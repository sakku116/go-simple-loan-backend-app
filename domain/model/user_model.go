package model

import (
	"backend/domain/enum"
	"backend/utils/helper"
	validator_util "backend/utils/validator/user"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UUID          string         `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	Username      string         `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password      string         `gorm:"type:varchar(255);not null" json:"-"`
	Email         string         `gorm:"type:varchar(255);email not null" json:"email"`
	Role          enum.UserRole  `gorm:"type:varchar(255);not null" json:"role"`
	Fullname      string         `gorm:"type:varchar(255);not null" json:"fullname"`
	Legalname     string         `gorm:"type:varchar(255);not null" json:"legalname"`
	NIK           string         `gorm:"type:varchar(255);not null" json:"nik"`
	Birthplace    string         `gorm:"type:varchar(255);not null" json:"birthplace"`
	Birthdate     string         `gorm:"type:varchar(255);not null" json:"birthdate"` // DD-MM-YYYY
	CurrentSalary int64          `gorm:"type:bigint;not null" json:"current_salary"`

	// these are required for requesting for loan
	KtpPhoto  *string `gorm:"type:varchar(255);null" json:"ktp_photo"`
	FacePhoto *string `gorm:"type:varchar(255);null" json:"face_photo"`

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;references:ID;" json:"-"`
}

func (u *User) GetProps() ModelProps {
	return ModelProps{
		MinioBucketName: "users",
		QueriableFields: helper.GetStructAttributesJson(u, []string{ // exclude by field names
			"ID", "UUID", "CreatedAt", "UpdatedAt", "DeletedAt",
			"Password", "KtpPhoto", "FacePhoto", "RefreshTokens",
		}, nil),
		SortableFields: []string{
			"updated_at", "created_at", "username", "fullname", "legalname",
		},
		DefaultSortableField: "updated_at",
	}
}

func (u *User) Validate() (err error) {
	// role
	isValid := u.Role.IsValid()
	if !isValid {
		return errors.New("user validation error: invalid role")
	}

	// username
	err = validator_util.ValidateUsername(u.Username)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// email
	err = validator_util.ValidateEmail(u.Email)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// password
	err = validator_util.ValidatePassword(u.Password)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// nik
	err = validator_util.ValidateNIK(u.NIK)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// fullname
	err = validator_util.ValidateFullname(u.Fullname)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// birthdate
	err = validator_util.ValidateBirthdate(u.Birthdate)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// birthplace
	err = validator_util.ValidateBirthplace(u.Birthplace)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	// legalname
	err = validator_util.ValidateLegalname(u.Legalname)
	if err != nil {
		return errors.New("user validation error: " + err.Error())
	}

	return nil
}
