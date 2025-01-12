package repository

import (
	"backend/domain/model"
	"errors"

	"gorm.io/gorm"
)

type LoanRepo struct {
	db *gorm.DB
}

type ILoanRepo interface {
	Create(loan *model.Loan) (*model.Loan, error)
	GetPaidListByUserID(id uint) ([]model.Loan, error)
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

func (repoo *LoanRepo) GetPaidListByUserID(id uint) ([]model.Loan, error) {
	var loans []model.Loan

	if err := repoo.db.Where("user_id = ?", id).Find(&loans).Error; err != nil {
		return nil, errors.New("failed to get loan: " + err.Error())
	}

	return loans, nil
}
