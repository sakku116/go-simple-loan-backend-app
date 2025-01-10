package model

import (
	"backend/domain/enum"
	validator_util "backend/utils/validator/user"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string        `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	Username string        `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password string        `gorm:"type:varchar(255);not null" json:"-"`
	Email    string        `gorm:"type:varchar(255);email not null" json:"email"`
	Role     enum.UserRole `gorm:"type:varchar(255);not null" json:"role"`

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;references:ID;" json:"-"`
}

func (u *User) Validate() (err error) {
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

	return nil
}
