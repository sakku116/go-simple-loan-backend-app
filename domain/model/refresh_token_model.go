package model

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	UUID      string     `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	UserUUID  string     `gorm:"type:varchar(36);not null" json:"user_uuid"`
	Token     string     `gorm:"type:varchar(255);unique;not null" json:"token"`
	UsedAt    *time.Time `gorm:"null" json:"used_at"`
	ExpiredAt *time.Time `gorm:"null" json:"expired_at"`
	Invalid   bool       `gorm:"type:bool;default:false" json:"invalid"`

	User User `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE;"`
}
