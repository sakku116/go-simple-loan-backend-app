package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	UserID   uint   `gorm:"type:uint;not null" json:"user_id"`
	UserUUID string `gorm:"type:varchar(36);not null" json:"user_uuid"`
	LoanID   uint   `gorm:"type:uint;not null" json:"loan_id"`
	LoanUUID string `gorm:"type:varchar(36);not null" json:"loan_uuid"`

	AssetName      string `gorm:"type:varchar(255);not null" json:"asset_name"`
	RefNumber      int64  `gorm:"type:bigint;not null" json:"ref_number"`
	Amount         int64  `gorm:"type:bigint;not null" json:"amount"`
	AdminFeeAmount int64  `gorm:"type:bigint;null" json:"adminFee"`
	RateAmount     int64  `gorm:"type:bigint;not null" json:"rate_amount"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Loan Loan `gorm:"foreignKey:LoanID;references:ID;constraint:OnDelete:CASCADE;"`
}
