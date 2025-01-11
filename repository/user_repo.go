package repository

import (
	"backend/domain/dto"
	"backend/domain/model"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type IUserRepo interface {
	Create(user *model.User) (*model.User, error)
	GetByUUID(uuid string) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByNIK(nik string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id string) error
	GetList(
		params dto.UserRepo_GetListParams,
	) ([]model.User, int64, error)
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Create(user *model.User) (*model.User, error) {
	err := repo.db.Create(user).Error
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	return user, err
}

func (repo *UserRepo) GetByUUID(uuid string) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "uuid = ?", uuid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("not found")
		}
		return nil, errors.New("failed to get: " + err.Error())
	}
	return &user, nil
}

func (repo *UserRepo) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "uuid = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("not found")
		}
		return nil, errors.New("failed to get: " + err.Error())
	}
	return &user, nil
}

func (repo *UserRepo) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetByNIK(nik string) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "nik = ?", nik).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) Update(user *model.User) (*model.User, error) {
	err := repo.db.Save(user).Error
	return user, err
}

func (repo *UserRepo) Delete(id string) error {
	err := repo.db.Delete(&model.User{}, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("not found")
		}
		return errors.New("failed to delete: " + err.Error())
	}
	return nil
}

func (repo *UserRepo) GetList(
	params dto.UserRepo_GetListParams,
) ([]model.User, int64, error) {
	var result []model.User
	var totalData int64

	// validate param
	err := params.Validate()
	if err != nil {
		return result, totalData, err
	}

	tx := repo.db.Model(&result)

	// filtering
	if params.Query != nil {
		if params.QueryBy != nil {
			tx = tx.Where(fmt.Sprintf("%s LIKE ?", *params.QueryBy), "%"+*params.Query+"%")
		} else {
			// filter by all queriable fields
			conditions := ""
			conditionValues := []interface{}{}
			tmp := model.User{}
			queriableFields := tmp.GetProps().QueriableFields
			for _, field := range queriableFields {
				logger.Debugf("field: %s", field)
				if field == "" {
					logger.Debugf("skipping empty field")
					continue
				}
				conditions += fmt.Sprintf(
					`%s LIKE ? OR `,
					field,
				)
				conditionValues = append(conditionValues, "%"+*params.Query+"%")
			}
			logger.Debugf("conditionValues: %v", conditionValues)
			conditions = strings.TrimSuffix(conditions, " OR ")
			tx = tx.Where(
				conditions,
				conditionValues...,
			)
		}
	}

	// get count if needed
	if params.DoCount {
		err = tx.Count(&totalData).Error
		if err != nil {
			return nil, totalData, errors.New("failed to count: " + err.Error())
		}
	}

	// sorting
	if params.SortOrder != nil && params.SortBy != nil {
		tx = tx.Order(fmt.Sprintf("%s %s", *params.SortBy, *params.SortOrder))
	}

	// pagination
	if params.Page != nil && params.Limit != nil {
		tx = tx.Offset((*params.Page - 1) * *params.Limit).Limit(*params.Limit)
	}

	err = tx.Find(&result).Error
	if err != nil {
		return nil, totalData, errors.New("failed to get: " + err.Error())
	}

	return result, totalData, nil
}
