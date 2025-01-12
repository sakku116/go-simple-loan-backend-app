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
	UpdateLoanStatus(
		loanUUID string,
		payload dto.UpdateLoanStatusReq,
	) (*dto.UpdateLoanStatusRespData, error)
	GetLoanList(
		params dto.GetLoanListReq,
	) (*dto.GetLoanListRespData, error)
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

	// validate total amount by current user loan limit
	if user.CurrentLimit == 0 {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  "user has no current limit",
		}
	}
	existingUnpaidLoan, err := ucase.loanRepo.GetUnPaidListByUserID(user.ID)
	if err != nil {
		logger.Debugf("failed to get existing closed loans: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}
	var currentUsedLimit float64
	for _, loan := range existingUnpaidLoan {
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
		BaseLoanResponse:      newLoan.ToBaseResponse(),
		CurrentLimitRemaining: user.CurrentLimit - currentUsedLimit,
	}, nil
}

func (ucase *LoanUcase) UpdateLoanStatus(
	loanUUID string,
	payload dto.UpdateLoanStatusReq,
) (*dto.UpdateLoanStatusRespData, error) {
	err := payload.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}

	// find loan
	loan, err := ucase.loanRepo.GetByUUID(loanUUID)
	if err != nil {
		logger.Debugf("failed to get loan: %v", err)
		if strings.Contains(err.Error(), "not found") {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "loan not found",
				Detail:   err.Error(),
			}
		}
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// update loan
	loan.Status = payload.Status
	_, err = ucase.loanRepo.Update(loan)
	if err != nil {
		logger.Debugf("failed to update loan: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	return &dto.UpdateLoanStatusRespData{
		Status: loan.Status,
	}, nil
}

func (ucase *LoanUcase) GetLoanList(
	params dto.GetLoanListReq,
) (*dto.GetLoanListRespData, error) {
	err := params.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  "bad request",
			Detail:   err.Error(),
		}
	}

	// prepare dto
	repoDto := dto.LoanRepo_GetListParams{
		UserUUID:  params.UserUUID,
		Status:    params.Status,
		Query:     params.Query,
		QueryBy:   params.QueryBy,
		Page:      &params.Page,
		Limit:     &params.Limit,
		SortOrder: params.SortOrder,
		SortBy:    params.SortBy,
		DoCount:   true,
	}

	data, totalData, err := ucase.loanRepo.GetList(repoDto)
	if err != nil {
		logger.Debugf("failed to get loan list: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// response
	resp := &dto.GetLoanListRespData{}
	resp.SetPagination(totalData, params.Page, params.Limit)
	for _, loan := range data {
		resp.Data = append(resp.Data, loan.ToBaseResponse())
	}

	return resp, nil
}
