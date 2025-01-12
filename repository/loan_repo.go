package repository

import (
	"backend/domain/dto"
	"backend/domain/model"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type LoanRepo struct {
	db *gorm.DB
}

type ILoanRepo interface {
	Create(loan *model.Loan) (*model.Loan, error)
	GetUnPaidListByUserID(id uint) ([]model.Loan, error)
	Update(loan *model.Loan) (*model.Loan, error)
	GetByUUID(uuid string) (*model.Loan, error)
	GetList(
		params dto.LoanRepo_GetListParams,
	) ([]model.Loan, int64, error)
}

func NewLoanRepo(db *gorm.DB) ILoanRepo {
	return &LoanRepo{db: db}
}

func (repo *LoanRepo) Create(loan *model.Loan) (*model.Loan, error) {
	err := repo.db.Create(loan).Error
	if err != nil {
		return nil, errors.New("failed to create loan")
	}
	return loan, err
}

func (repoo *LoanRepo) GetUnPaidListByUserID(id uint) ([]model.Loan, error) {
	var loans []model.Loan

	if err := repoo.db.Where("user_id = ?", id).Where("status <> ?", "paid").Find(&loans).Error; err != nil {
		return nil, errors.New("failed to get loan: " + err.Error())
	}

	return loans, nil
}

func (repo *LoanRepo) Update(loan *model.Loan) (*model.Loan, error) {
	err := repo.db.Save(loan).Error
	if err != nil {
		return nil, errors.New("failed to update loan")
	}
	return loan, nil
}

func (repo *LoanRepo) GetByUUID(uuid string) (*model.Loan, error) {
	var loan model.Loan
	if err := repo.db.First(&loan, "uuid = ?", uuid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("not found")
		}
		return nil, errors.New("failed to get: " + err.Error())
	}
	return &loan, nil
}

func (repo *LoanRepo) GetList(
	params dto.LoanRepo_GetListParams,
) ([]model.Loan, int64, error) {
	var result []model.Loan
	var totalData int64

	// validate param
	err := params.Validate()
	if err != nil {
		return result, totalData, err
	}

	tx := repo.db.Model(&result)

	// filtering
	if params.UserUUID != nil {
		tx = tx.Where("user_uuid = ?", *params.UserUUID)
	}
	if params.Status != nil {
		tx = tx.Where("status = ?", *params.Status)
	}
	if params.Query != nil {
		if params.QueryBy != nil {
			tx = tx.Where(fmt.Sprintf("%s LIKE ?", *params.QueryBy), "%"+*params.Query+"%")
		} else {
			// filter by all queriable fields
			conditions := ""
			conditionValues := []interface{}{}
			tmp := model.Loan{}
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
