package ucase

import (
	"backend/domain/dto"
	"backend/domain/enum"
	"backend/domain/model"
	"backend/repository"
	error_utils "backend/utils/error"
	"backend/utils/helper"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type LoanUcase struct {
	loanRepo repository.ILoanRepo
	userRepo repository.IUserRepo
}

type ILoanUcase interface {
	CreateNewLoan(
		userUUID string,
		payload dto.CreateNewLoanReq,
	) (*dto.CreateNewLoanRespData, error)
}

func NewLoanuCase(
	loanRepo repository.ILoanRepo,
	userRepo repository.IUserRepo,
) ILoanUcase {
	return &LoanUcase{
		loanRepo: loanRepo,
		userRepo: userRepo,
	}
}

func (ucase *LoanUcase) CreateNewLoan(
	userUUID string,
	payload dto.CreateNewLoanReq,
) (*dto.CreateNewLoanRespData, error) {
	err := payload.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}

	// find user
	user, err := ucase.userRepo.GetByUUID(userUUID)
	if err != nil {
		logger.Debugf("failed to get user: %v", err)
		if strings.Contains(err.Error(), "not found") {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// create new loan
	newLoan := &model.Loan{
		UUID:                   uuid.New().String(),
		UserID:                 user.ID,
		UserUUID:               user.UUID,
		RefNumber:              helper.TimeNowEpochUtc(),
		OTR:                    payload.OTR,
		AssetName:              payload.AssetName,
		TermMonths:             payload.TermMonths,
		Status:                 enum.LoanStatus_PENDING,
		AdminFeePercentage:     2,
		InterestRatePercentage: 10,
	}

	// interest rate calculation (change if needed)
	newLoan.InterestRate = float64(newLoan.OTR) * (newLoan.InterestRatePercentage / 100)

	// admin fee calculation (change if needed)
	newLoan.AdminFee = float64(newLoan.OTR) * (newLoan.AdminFeePercentage / 100)

	// calculate total amount
	newLoan.TotalAmount = newLoan.OTR + newLoan.InterestRate + newLoan.AdminFee

	// calculate installment amount
	newLoan.InstallmentAmount = (newLoan.OTR + newLoan.InterestRate + newLoan.AdminFee) / float64(newLoan.TermMonths)

	// pretend that loan status has been changed to APPROVED by admin or advanced system logic(separate logic if needed)
	newLoan.Status = enum.LoanStatus_APPROVED

	// validate OTR by current user loan limit
	if user.CurrentLimit == 0 {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  "user has no current limit",
		}
	}
	existingClosedLoans, err := ucase.loanRepo.GetPaidListByUserID(user.ID)
	if err != nil {
		logger.Debugf("failed to get existing closed loans: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}
	var currentUsedLimit float64
	for _, loan := range existingClosedLoans {
		currentUsedLimit += loan.TotalAmount
	}
	if newLoan.TotalAmount+currentUsedLimit > float64(user.CurrentLimit) {
		logger.Debugf("limit not enough")
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  fmt.Sprintf("limit not enough, current limit remaining: %v", user.CurrentLimit-currentUsedLimit),
		}
	}

	// save
	_, err = ucase.loanRepo.Create(newLoan)
	if err != nil {
		logger.Debugf("failed to create new loan: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	currentUsedLimit += newLoan.TotalAmount
	return &dto.CreateNewLoanRespData{
		BaseLoanResponse: newLoan.ToBaseResponse(user.CurrentLimit, currentUsedLimit),
	}, nil
}
