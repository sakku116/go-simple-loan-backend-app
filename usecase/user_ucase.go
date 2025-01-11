package ucase

import (
	"backend/domain/dto"
	"backend/domain/model"
	"backend/repository"
	bcrypt_util "backend/utils/bcrypt"
	error_utils "backend/utils/error"
	file_storage_util "backend/utils/file"
	"backend/utils/helper"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
)

type UserUcase struct {
	userRepo        repository.IUserRepo
	fileStorageUtil file_storage_util.IFileStorageUtil
}

type IUserUcase interface {
	GetByUUID(ctx context.Context, userUUID string) (*dto.GetUserByUUIDResp, error)
	CreateUser(
		ctx context.Context,
		payload dto.CreateUserReq,
	) (*dto.CreateUserRespData, error)
	UpdateUser(
		ctx context.Context,
		userUUID string,
		payload dto.UpdateUserReq,
	) (*dto.UpdateUserRespData, error)
	DeleteUser(
		ctx context.Context,
		userUUID string,
	) (*dto.DeleteUserRespData, error)
	UploadKtpPhoto(ctx context.Context, userUUID string, file *multipart.FileHeader) (*dto.UploadKtpPhotoRespData, error)
	UploadFacePhoto(ctx context.Context, userUUID string, file *multipart.FileHeader) (*dto.UploadFacePhotoRespData, error)
	GetUserList(
		ctx context.Context,
		params dto.GetUserListReq,
	) (*dto.GetUserListRespData, error)
}

func NewUserUcase(
	userRepo repository.IUserRepo,
	fileStorageUtil file_storage_util.IFileStorageUtil,
) IUserUcase {
	return &UserUcase{
		userRepo:        userRepo,
		fileStorageUtil: fileStorageUtil,
	}
}

func (ucase *UserUcase) GetByUUID(ctx context.Context, userUUID string) (*dto.GetUserByUUIDResp, error) {
	user, err := ucase.userRepo.GetByUUID(userUUID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, err
	}

	return &dto.GetUserByUUIDResp{
		BaseUserResponse: user.ToBaseResponse(ctx, ucase.fileStorageUtil),
	}, nil
}

func (ucase *UserUcase) CreateUser(
	ctx context.Context,
	payload dto.CreateUserReq,
) (*dto.CreateUserRespData, error) {
	// validate input
	err := payload.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}

	// check if user exists
	user, _ := ucase.userRepo.GetByEmail(payload.Email)
	logger.Debugf("user by email: %v", user)
	if user != nil {
		logger.Errorf("user with email %s already exists", payload.Email)
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  fmt.Sprintf("user with email %s already exists", payload.Email),
		}
	}

	user, _ = ucase.userRepo.GetByUsername(payload.Username)
	if user != nil {
		logger.Errorf("user with username %s already exists", payload.Username)
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  fmt.Sprintf("user with username %s already exists", payload.Username),
		}
	}

	// create password
	password, err := bcrypt_util.Hash(payload.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err)
		return nil, err
	}

	// create user
	user = &model.User{
		UUID:          uuid.New().String(),
		Username:      payload.Username,
		Password:      password,
		Email:         payload.Email,
		Role:          payload.Role,
		Fullname:      payload.Fullname,
		Legalname:     payload.Legalname,
		NIK:           payload.NIK,
		Birthplace:    payload.Birthplace,
		Birthdate:     payload.Birthdate,
		CurrentSalary: payload.CurrentSalary,
	}
	err = user.Validate()
	if err != nil {
		return nil, err
	}

	user, err = ucase.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	userResp := user.ToBaseResponse(ctx, ucase.fileStorageUtil)

	return &dto.CreateUserRespData{
		BaseUserResponse: userResp,
	}, nil
}

func (ucase *UserUcase) UpdateUser(
	ctx context.Context,
	userUUID string,
	payload dto.UpdateUserReq,
) (*dto.UpdateUserRespData, error) {
	// validate input
	err := payload.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}

	// get existing user
	user, err := ucase.userRepo.GetByUUID(userUUID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, err
	}

	// update user obj
	if payload.Username != nil {
		user.Username = *payload.Username
	}
	if payload.Email != nil {
		user.Email = *payload.Email
	}
	if payload.Password != nil {
		password, err := bcrypt_util.Hash(*payload.Password)
		if err != nil {
			logger.Errorf("error hashing password: %v", err)
			return nil, err
		}
		user.Password = password
	}
	if payload.Role != nil {
		user.Role = *payload.Role
	}
	if payload.Fullname != nil {
		user.Fullname = *payload.Fullname
	}
	if payload.Legalname != nil {
		user.Legalname = *payload.Legalname
	}
	if payload.NIK != nil {
		user.NIK = *payload.NIK
	}
	if payload.Birthplace != nil {
		user.Birthplace = *payload.Birthplace
	}
	if payload.Birthdate != nil {
		user.Birthdate = *payload.Birthdate
	}
	if payload.CurrentSalary != nil {
		user.CurrentSalary = *payload.CurrentSalary
	}

	// update user
	user, err = ucase.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserRespData{
		BaseUserResponse: user.ToBaseResponse(ctx, ucase.fileStorageUtil),
	}, nil
}

func (ucase *UserUcase) DeleteUser(
	ctx context.Context,
	userUUID string,
) (*dto.DeleteUserRespData, error) {
	// find user
	user, err := ucase.userRepo.GetByUUID(userUUID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, err
	}

	// delete user
	err = ucase.userRepo.Delete(user.UUID)
	if err != nil {
		if err.Error() == "not found" {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, err
	}

	return &dto.DeleteUserRespData{
		BaseUserResponse: user.ToBaseResponse(ctx, ucase.fileStorageUtil),
	}, nil
}

func (u *UserUcase) UploadKtpPhoto(ctx context.Context, userUUID string, file *multipart.FileHeader) (*dto.UploadKtpPhotoRespData, error) {
	// find user
	user, err := u.userRepo.GetByUUID(userUUID)
	if err != nil {
		logger.Debugf("failed to get user: %v", err)
		if err.Error() == "not found" {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// upload file
	tmp := model.User{}
	filename := fmt.Sprintf("%s-%s", uuid.New().String(), file.Filename)
	filename, err = u.fileStorageUtil.Upload(ctx, file, filename, tmp.GetProps().BucketName)
	if err != nil {
		logger.Debugf("failed to upload KTP photo: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// update user
	user.KtpPhoto = &filename
	user.UpdatedAt = helper.TimeNowUTC()
	user, err = u.userRepo.Update(user)
	if err != nil {
		logger.Debugf("failed to update user: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// get url
	url, err := u.fileStorageUtil.GetUrl(ctx, filename, tmp.GetProps().BucketName, false)
	if err != nil {
		logger.Debugf("failed to get url: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	return &dto.UploadKtpPhotoRespData{
		KtpPhoto: url,
	}, nil
}

func (u *UserUcase) UploadFacePhoto(ctx context.Context, userUUID string, file *multipart.FileHeader) (*dto.UploadFacePhotoRespData, error) {
	// find user
	user, err := u.userRepo.GetByUUID(userUUID)
	if err != nil {
		logger.Debugf("failed to get user: %v", err)
		if err.Error() == "not found" {
			return nil, &error_utils.CustomErr{
				HttpCode: 404,
				Message:  "user not found",
				Detail:   err.Error(),
			}
		}
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// upload file
	tmp := model.User{}
	filename := fmt.Sprintf("%s-%s", uuid.New().String(), file.Filename)
	filename, err = u.fileStorageUtil.Upload(ctx, file, filename, tmp.GetProps().BucketName)
	if err != nil {
		logger.Debugf("failed to upload KTP photo: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// update user
	user.FacePhoto = &filename
	user.UpdatedAt = helper.TimeNowUTC()
	user, err = u.userRepo.Update(user)
	if err != nil {
		logger.Debugf("failed to update user: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// get url
	url, err := u.fileStorageUtil.GetUrl(ctx, filename, tmp.GetProps().BucketName, false)
	if err != nil {
		logger.Debugf("failed to get url: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	return &dto.UploadFacePhotoRespData{
		FacePhoto: url,
	}, nil
}

func (u *UserUcase) GetUserList(
	ctx context.Context,
	params dto.GetUserListReq,
) (*dto.GetUserListRespData, error) {
	// prepare dto
	repoDto := dto.UserRepo_GetListParams{
		Query:     params.Query,
		Page:      &params.Page,
		Limit:     &params.Limit,
		SortOrder: params.SortOrder,
		SortBy:    params.SortBy,
		DoCount:   true,
	}
	if params.QueryBy != nil && *params.QueryBy != "" {
		repoDto.QueryBy = params.QueryBy
	} else {
		repoDto.QueryBy = nil
	}

	data, totalData, err := u.userRepo.GetList(repoDto)
	if err != nil {
		logger.Debugf("failed to get user list: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "internal server error",
			Detail:   err.Error(),
		}
	}

	// response
	resp := &dto.GetUserListRespData{}
	resp.SetPagination(totalData, params.Page, params.Limit)
	if data != nil {
		for _, v := range data {
			resp.Data = append(resp.Data, v.ToBaseResponse(ctx, u.fileStorageUtil))
		}
	}
	return resp, nil
}
