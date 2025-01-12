package model

import (
	"backend/domain/enum"
	file_storage_util "backend/utils/file"
	"backend/utils/helper"
	validator_util "backend/utils/validator/user"
	"context"
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID          string        `gorm:"type:varchar(36);unique;not null" json:"uuid"`
	Username      string        `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password      string        `gorm:"type:varchar(255);not null" json:"-"`
	Email         string        `gorm:"type:varchar(255);email not null" json:"email"`
	Role          enum.UserRole `gorm:"type:varchar(255);not null" json:"role"`
	Fullname      string        `gorm:"type:varchar(255);not null" json:"fullname"`
	Legalname     string        `gorm:"type:varchar(255);not null" json:"legalname"`
	NIK           string        `gorm:"type:varchar(255);not null" json:"nik"`
	Birthplace    string        `gorm:"type:varchar(255);not null" json:"birthplace"`
	Birthdate     string        `gorm:"type:varchar(255);not null" json:"birthdate"` // DD-MM-YYYY
	CurrentSalary float64       `gorm:"type:float;not null" json:"current_salary"`
	CurrentLimit  float64       `gorm:"type:float;not null" json:"current_limit"`

	// these are required for requesting for loan
	KtpPhoto  *string `gorm:"type:varchar(255);null" json:"ktp_photo"`
	FacePhoto *string `gorm:"type:varchar(255);null" json:"face_photo"`

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;references:ID;" json:"-"`
}

type BaseUserResponse struct {
	UUID          string  `json:"uuid"`
	Username      string  `json:"username"`
	Email         string  `json:"email"`
	Role          string  `json:"role"`
	Fullname      string  `json:"fullname"`
	Legalname     string  `json:"legalname"`
	NIK           string  `json:"nik"`
	Birthplace    string  `json:"birthplace"`
	Birthdate     string  `json:"birthdate"`
	CurrentSalary float64 `json:"current_salary"`
	CurrentLimit  float64 `json:"current_limit"`
	KtpPhoto      *string `json:"ktp_photo"`
	FacePhoto     *string `json:"face_photo"`
}

func (u *User) ToBaseResponse(ctx context.Context, fileStorageUtil file_storage_util.IFileStorageUtil) BaseUserResponse {
	logger.Debugf(helper.PrettyJson(u))
	bucketName := u.GetProps().BucketName
	// urlize file fields
	if u.KtpPhoto != nil {
		tmp, _ := fileStorageUtil.GetUrl(ctx, *u.KtpPhoto, bucketName, false)
		if tmp != "" {
			u.KtpPhoto = &tmp
		}
	}
	if u.FacePhoto != nil {
		tmp, _ := fileStorageUtil.GetUrl(ctx, *u.FacePhoto, bucketName, false)
		if tmp != "" {
			u.FacePhoto = &tmp
		}
	}
	return BaseUserResponse{
		UUID:          u.UUID,
		Username:      u.Username,
		Email:         u.Email,
		Role:          u.Role.String(),
		Fullname:      u.Fullname,
		Legalname:     u.Legalname,
		NIK:           u.NIK,
		Birthplace:    u.Birthplace,
		Birthdate:     u.Birthdate,
		CurrentSalary: u.CurrentSalary,
		KtpPhoto:      u.KtpPhoto,
		FacePhoto:     u.FacePhoto,
		CurrentLimit:  u.CurrentLimit,
	}
}

func (u *User) GetProps() ModelProps {
	return ModelProps{
		BucketName: "users",
		QueriableFields: helper.GetStructAttributesJson(u,
			[]string{ // exclude by field names
				"ID", "UUID", "CreatedAt", "UpdatedAt", "DeletedAt",
				"Password", "KtpPhoto", "FacePhoto", "RefreshTokens",
			},
			[]string{
				"-",
			},
		),
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
