package dto

import (
	"backend/domain/enum"
	"backend/domain/model"
	"errors"
	"fmt"
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
	CurrentLimitRemaining float64 `json:"current_limit_remaining"`
}

type GetLoanListReq struct {
	UserUUID *string          `form:"user_uuid" binding:"omitempty"`
	Status   *enum.LoanStatus `form:"status" binding:"omitempty"`

	Query     *string         `form:"query" binding:"omitempty"`
	QueryBy   *string         `form:"query_by" binding:"omitempty,oneof=asset_name ref_number"` // leave empty to query by all
	Page      int             `form:"page" default:"1"`
	Limit     int             `form:"limit" default:"10"`
	SortOrder *enum.SortOrder `form:"sort_order" binding:"omitempty,oneof=asc desc" default:"desc"`
	SortBy    *string         `form:"sort_by" binding:"omitempty,oneof=updated_at" default:"updated_at"`
}

func (params *GetLoanListReq) Validate() error {
	if params.SortOrder != nil && !(*params.SortOrder).IsValid() {
		return fmt.Errorf("invalid sort_order")
	}

	if params.QueryBy != nil && *params.QueryBy == "" {
		params.QueryBy = nil
	}

	return nil
}

type GetLoanListRespData struct {
	BasePaginationRespData
	Data []model.BaseLoanResponse `json:"data"`
}

type LoanRepo_GetListParams struct {
	UserUUID *string
	Status   *enum.LoanStatus

	Query     *string
	QueryBy   *string // leave empty to query by all
	Page      *int
	Limit     *int
	SortOrder *enum.SortOrder
	SortBy    *string
	DoCount   bool
}

func (params *LoanRepo_GetListParams) Validate() error {
	if params.SortOrder != nil && !(*params.SortOrder).IsValid() {
		return fmt.Errorf("invalid sort_order")
	}

	tmp := model.Loan{}
	if params.QueryBy != nil {
		queriableFields := tmp.GetProps().QueriableFields
		contain := false
		for _, field := range queriableFields {
			if *params.QueryBy == field {
				contain = true
				break
			}
		}
		if !contain {
			return fmt.Errorf("invalid query_by")
		}
	}

	if params.SortBy != nil {
		sortableFields := tmp.GetProps().SortableFields
		contain := false
		for _, field := range sortableFields {
			if *params.SortBy == field {
				contain = true
				break
			}
		}
		if !contain {
			return fmt.Errorf("invalid sort_by")
		}
	}

	return nil
}

type UpdateLoanStatusReq struct {
	Status enum.LoanStatus `json:"status" binding:"required"`
}

func (req *UpdateLoanStatusReq) Validate() error {
	if !req.Status.IsValid() {
		return fmt.Errorf("invalid status, status must be %v", enum.ValidLoanStatus)
	}

	return nil
}

type UpdateLoanStatusRespData struct {
	Status enum.LoanStatus `json:"status"`
}
