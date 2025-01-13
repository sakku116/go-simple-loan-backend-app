package model

import (
	"backend/domain/enum"
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	UserID   uint   `gorm:"type:uint;not null" json:"user_id"`
	UserUUID string `gorm:"type:varchar(36);not null" json:"user_uuid"`

	AssetName              string              `gorm:"type:varchar(255);not null" json:"asset_name"`
	RefNumber              int64               `gorm:"type:bigint;not null" json:"ref_number"`
	OTR                    float64             `gorm:"type:float;not null" json:"otr"`
	InterestRatePercentage float64             `gorm:"type:float;not null" json:"interest_rate_percentage" default:"10"`
	InterestRate           float64             `gorm:"type:float;not null" json:"interest_rate"`
	AdminFeePercentage     float64             `gorm:"type:float;not null" json:"admin_fee_percentage" default:"2"`
	AdminFee               float64             `gorm:"type:float;not null" json:"admin_fee"`
	InstallmentAmount      float64             `gorm:"type:float;not null" json:"installment_amount"`
	TotalAmount            float64             `gorm:"type:float;not null" json:"total_amount"`
	TermMonths             enum.LoanTermMonths `gorm:"type:int;null" json:"term_months"`
	Status                 enum.LoanStatus     `gorm:"type:varchar(255);not null" json:"status" default:"PENDING"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

func (u *Loan) GetProps() ModelProps {
	return ModelProps{
		QueriableFields: []string{
			"asset_name", "ref_number",
		},
		SortableFields: []string{
			"updated_at", "created_at", "asset_name",
		},
		DefaultSortableField: "updated_at",
	}
}

type BaseLoanResponse struct {
	UUID                   string              `json:"uuid"`
	CreatedAt              time.Time           `json:"created_at"`
	UpdatedAt              time.Time           `json:"updated_at"`
	UserUUID               string              `json:"user_uuid"`
	AssetName              string              `json:"asset_name"`
	RefNumber              int64               `json:"ref_number"`
	OTR                    float64             `json:"otr"`
	InterestRatePercentage float64             `json:"interest_rate_percentage"`
	InterestRate           float64             `json:"interest_rate"`
	AdminFeePercentage     float64             `json:"admin_fee_percentage"`
	AdminFee               float64             `json:"admin_fee"`
	InstallmentAmount      float64             `json:"installment_amount"`
	TotalAmount            float64             `json:"total_amount"`
	TermMonths             enum.LoanTermMonths `json:"term_months"`
	Status                 enum.LoanStatus     `json:"status"`
}

func (loan *Loan) ToBaseResponse() BaseLoanResponse {
	data := BaseLoanResponse{
		UUID:                   loan.UUID,
		CreatedAt:              loan.CreatedAt,
		UpdatedAt:              loan.UpdatedAt,
		UserUUID:               loan.UserUUID,
		AssetName:              loan.AssetName,
		RefNumber:              loan.RefNumber,
		OTR:                    loan.OTR,
		InterestRatePercentage: loan.InterestRatePercentage,
		InterestRate:           loan.InterestRate,
		AdminFeePercentage:     loan.AdminFeePercentage,
		AdminFee:               loan.AdminFee,
		InstallmentAmount:      loan.InstallmentAmount,
		TotalAmount:            loan.TotalAmount,
		TermMonths:             loan.TermMonths,
		Status:                 loan.Status,
	}
	return data
}
