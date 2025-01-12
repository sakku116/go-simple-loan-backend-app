package dto

import (
	"backend/domain/enum"
	"backend/domain/model"
	validator_util "backend/utils/validator/user"
	"errors"
	"fmt"
	"mime/multipart"
)

type GetUserByUUIDResp struct {
	model.BaseUserResponse
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
	CurrentSalary float64       `json:"current_salary" validate:"required"`
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
	model.BaseUserResponse
}

type UpdateUserReq struct {
	Username      *string        `json:"username"`
	Email         *string        `json:"email"`
	Password      *string        `json:"password"`
	Role          *enum.UserRole `json:"role"`
	Fullname      *string        `json:"fullname"`
	Legalname     *string        `json:"legalname"`
	NIK           *string        `json:"nik"`
	Birthplace    *string        `json:"birthplace"`
	Birthdate     *string        `json:"birthdate"` // DD-MM-YYYY
	CurrentSalary *float64       `json:"current_salary"`
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

	if req.Role != nil {
		isValid := (*req.Role).IsValid()
		if !isValid {
			return errors.New("user validation error: invalid role")
		}
	}

	if req.Fullname != nil {
		err := validator_util.ValidateFullname(*req.Fullname)
		if err != nil {
			return err
		}
	}

	if req.Legalname != nil {
		err := validator_util.ValidateLegalname(*req.Legalname)
		if err != nil {
			return err
		}
	}

	if req.NIK != nil {
		err := validator_util.ValidateNIK(*req.NIK)
		if err != nil {
			return err
		}
	}

	if req.Birthplace != nil {
		err := validator_util.ValidateBirthplace(*req.Birthplace)
		if err != nil {
			return err
		}
	}

	if req.Birthdate != nil {
		err := validator_util.ValidateBirthdate(*req.Birthdate)
		if err != nil {
			return err
		}
	}

	return nil
}

type UpdateUserRespData struct {
	model.BaseUserResponse
}

type DeleteUserRespData struct {
	model.BaseUserResponse
}

type UploadKtpPhotoReq struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadKtpPhotoRespData struct {
	KtpPhoto string `json:"ktp_photo"`
}

type UploadFacePhotoReq struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadFacePhotoRespData struct {
	FacePhoto string `json:"face_photo"`
}

type UserRepo_GetListParams struct {
	Query     *string
	QueryBy   *string // leave empty to query by all
	Page      *int
	Limit     *int
	SortOrder *enum.SortOrder
	SortBy    *string
	DoCount   bool
}

func (params *UserRepo_GetListParams) Validate() error {
	tmp := model.User{}
	if params.QueryBy != nil {
		queriableFields := tmp.GetProps().QueriableFields
		contain := false
		for _, field := range queriableFields {
			if *params.QueryBy == field {
				contain = true
				break
			}
		}
		if !contain {
			return fmt.Errorf("invalid query_by")
		}
	}

	if params.SortBy != nil {
		sortableFields := tmp.GetProps().SortableFields
		contain := false
		for _, field := range sortableFields {
			if *params.SortBy == field {
				contain = true
				break
			}
		}
		if !contain {
			return fmt.Errorf("invalid sort_by")
		}
	}

	return nil
}

type GetUserListReq struct {
	Query     *string         `form:"query" binding:"omitempty"`
	QueryBy   *string         `form:"query_by" binding:"omitempty,oneof=username email nik fullname legalname role"` // leave empty
	Page      int             `form:"page" default:"1"`
	Limit     int             `form:"limit" default:"10"`
	SortOrder *enum.SortOrder `form:"sort_order" binding:"omitempty,oneof=asc desc" default:"desc"`
	SortBy    *string         `form:"sort_by" binding:"oneof=updated_at username email nik fullname legalname role" default:"updated_at"`
}

type GetUserListRespData struct {
	BasePaginationRespData
	Data []model.BaseUserResponse `json:"data"`
}

type UpdateCurrentLimitReq struct {
	CurrentLimit float64 `json:"current_limit" binding:"required"`
}

type UpdateCurrentLimitRespData struct {
	CurrentLimit float64 `json:"current_limit"`
}
