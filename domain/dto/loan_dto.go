package dto

import (
	"backend/domain/enum"
	"backend/domain/model"
	"errors"
)

type CreateNewLoanReq struct {
	AssetName  string              `json:"asset_name" binding:"required"`
	OTR        float64             `json:"otr" binding:"required"`
	TermMonths enum.LoanTermMonths `json:"term_months" binding:"required"`
}

func (r *CreateNewLoanReq) Validate() error {
	valid := r.TermMonths.IsValid()
	if !valid {
		return errors.New("invalid term months")
	}

	return nil
}

type CreateNewLoanRespData struct {
	model.BaseLoanResponse
}
