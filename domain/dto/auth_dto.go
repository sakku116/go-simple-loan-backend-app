package dto

import (
	"backend/domain/enum"
	"backend/domain/model"
	validator_util "backend/utils/validator/user"
	"strings"
)

type CurrentUser struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type RegisterUserReq struct {
	Email         string  `json:"email" validate:"required,email"`
	Username      string  `json:"username" validate:"required"`
	Password      string  `json:"password" validate:"required"`
	Fullname      string  `json:"fullname" validate:"required"`
	Legalname     string  `json:"legalname" validate:"required"`
	NIK           string  `json:"nik" validate:"required"`
	Birthplace    string  `json:"birthplace" validate:"required"`
	Birthdate     string  `json:"birthdate" validate:"required"` // DD-MM-YYYY
	CurrentSalary float64 `json:"current_salary" validate:"required"`
}

func (req *RegisterUserReq) Validate() error {
	tmp := model.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   req.Password,
		Role:       enum.UserRole("user"),
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

type RegisterUserRespData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginDevReq struct {
	Username string `form:"username" validate:"required"` // username or email, but swagger oauth2password need username field
	Password string `form:"password" validate:"required"`
}

func (req *LoginDevReq) Validate() error {
	if strings.Contains(req.Username, "@") {
		err := validator_util.ValidateEmail(req.Username)
		if err != nil {
			return err
		}
	} else {
		err := validator_util.ValidateUsername(req.Username)
		if err != nil {
			return err
		}
	}
	err := validator_util.ValidatePassword(req.Password)
	if err != nil {
		return err
	}

	return nil
}

type LoginReq struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

func (req *LoginReq) Validate() error {
	if strings.Contains(req.UsernameOrEmail, "@") {
		err := validator_util.ValidateEmail(req.UsernameOrEmail)
		if err != nil {
			return err
		}
	} else {
		err := validator_util.ValidateUsername(req.UsernameOrEmail)
		if err != nil {
			return err
		}
	}
	err := validator_util.ValidatePassword(req.Password)
	if err != nil {
		return err
	}

	return nil
}

type LoginRespData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginDevResp struct {
	BaseJSONResp
	AccessToken string `json:"access_token"`
}

type CheckTokenReq struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type CheckTokenRespData struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenRespData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
