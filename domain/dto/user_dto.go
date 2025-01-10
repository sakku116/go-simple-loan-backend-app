package dto

import (
	"backend/domain/enum"
	"backend/domain/model"
	validator_util "backend/utils/validator/user"
	"time"
)

type GetUserByUUIDResp struct {
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

type CreateUserReq struct {
	Username      string        `json:"username" binding:"required"`
	Email         string        `json:"email" binding:"required,email"`
	Password      string        `json:"password" binding:"required"`
	Role          enum.UserRole `json:"role" binding:"required,oneof=admin user"`
	Fullname      string        `json:"fullname" validate:"required"`
	Legalname     string        `json:"legalname" validate:"required"`
	NIK           string        `json:"nik" validate:"required"`
	Birthplace    string        `json:"birthplace" validate:"required"`
	Birthdate     string        `json:"birthdate" validate:"required"` // DD-MM-YYYY
	CurrentSalary int64         `json:"current_salary" validate:"required"`
}

func (req *CreateUserReq) Validate() error {
	tmp := model.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   req.Password,
		Role:       req.Role,
		Fullname:   req.Fullname,
		NIK:        req.NIK,
		Legalname:  req.Legalname,
		Birthplace: req.Birthplace,
		Birthdate:  req.Birthdate,
	}

	err := tmp.Validate()
	if err != nil {
		return err
	}

	return nil
}

type CreateUserRespData struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserReq struct {
	Username *string        `json:"username"`
	Email    *string        `json:"email"`
	Password *string        `json:"password"`
	Role     *enum.UserRole `json:"role" binding:"oneof=admin user"`
}

func (req *UpdateUserReq) Validate() error {
	if req.Username != nil {
		err := validator_util.ValidateUsername(*req.Username)
		if err != nil {
			return err
		}
	}

	if req.Email != nil {
		err := validator_util.ValidateEmail(*req.Email)
		if err != nil {
			return err
		}
	}

	if req.Password != nil {
		err := validator_util.ValidatePassword(*req.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

type UpdateUserRespData struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserRespData struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
