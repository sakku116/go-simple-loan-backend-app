package model

import (
	"backend/domain/enum"

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
	InterestRatePercentage float64             `gorm:"type:float;null" json:"interest_rate_percentage" default:"10"`
	InterestRate           float64             `gorm:"type:float;null" json:"interest_rate"`
	AdminFeePercentage     float64             `gorm:"type:float;null" json:"admin_fee_percentage" default:"2"`
	AdminFee               float64             `gorm:"type:float;null" json:"admin_fee"`
	InstallmentAmount      float64             `gorm:"type:float;null" json:"installment_amount"`
	TotalAmount            float64             `gorm:"type:float;null" json:"total_amount"`
	TermMonths             enum.LoanTermMonths `gorm:"type:int;null" json:"term_months"`
	Status                 enum.LoanStatus     `gorm:"type:varchar(255);not null" json:"status" default:"PENDING"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type BaseLoanResponse struct {
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
	CurrentLimitRemaining  float64             `json:"current_limit_remaining"`
}

func (loan *Loan) ToBaseResponse(userCurrentLimit, usedLimit float64) BaseLoanResponse {
	return BaseLoanResponse{
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
		CurrentLimitRemaining:  userCurrentLimit - usedLimit,
	}
}
