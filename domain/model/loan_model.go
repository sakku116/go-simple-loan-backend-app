package model

import (
	"backend/domain/enum"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	UUID         string              `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	UserID       uint                `gorm:"type:uint;not null" json:"user_id"`
	UserUUID     string              `gorm:"type:varchar(36);not null" json:"user_uuid"`
	Amount       int64               `gorm:"type:bigint;not null" json:"amount"`
	InterestRate float64             `gorm:"type:float;null" json:"interest_rate"` // flat rate
	TermMonths   enum.LoanTermMonths `gorm:"type:int;null" json:"term_months"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}
