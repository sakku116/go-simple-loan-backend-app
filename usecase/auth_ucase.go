package ucase

import (
	"backend/config"
	"backend/domain/dto"
	"backend/domain/enum"
	"backend/domain/model"
	"backend/repository"
	bcrypt_util "backend/utils/bcrypt"
	error_utils "backend/utils/error"
	"backend/utils/helper"
	jwt_util "backend/utils/jwt"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthUcase struct {
	userRepo         repository.IUserRepo
	refreshTokenRepo repository.IRefreshTokenRepo
}

type IAuthUcase interface {
	Register(ctx *gin.Context, payload dto.RegisterUserReq) (*dto.RegisterUserRespData, error)
	Login(payload dto.LoginReq) (*dto.LoginRespData, error)
	RefreshToken(payload dto.RefreshTokenReq) (*dto.RefreshTokenRespData, error)
	CheckToken(payload dto.CheckTokenReq) (*dto.CheckTokenRespData, error)
}

func NewAuthUcase(
	userRepo repository.IUserRepo,
	refreshTokenRepo repository.IRefreshTokenRepo,
) IAuthUcase {
	return &AuthUcase{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *AuthUcase) Register(ctx *gin.Context, payload dto.RegisterUserReq) (*dto.RegisterUserRespData, error) {
	// validate input
	err := payload.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}

	// check if user exists
	user, _ := s.userRepo.GetByEmail(payload.Email)
	logger.Debugf("user by email: %v", user)
	if user != nil {
		logger.Errorf("user with email %s already exists", payload.Email)
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  fmt.Sprintf("user with email %s already exists", payload.Email),
		}
	}

	user, _ = s.userRepo.GetByUsername(payload.Username)
	logger.Debugf("user by username: %v", user)
	if user != nil {
		logger.Errorf("user with username %s already exists", payload.Username)
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  fmt.Sprintf("user with username %s already exists", payload.Username),
		}
	}

	user, _ = s.userRepo.GetByNIK(payload.NIK)
	if user != nil {
		logger.Errorf("user with nik %s already exists", payload.NIK)
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  fmt.Sprintf("user with nik %s already exists", payload.NIK),
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
		Role:          enum.UserRole_User,
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

	user, err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// generate token
	token, err := jwt_util.GenerateJwtToken(user, config.Envs.JWT_SECRET_KEY, config.Envs.JWT_EXP_MINS, nil)
	if err != nil {
		logger.Errorf("error generating token: %v", err)
		return nil, err
	}

	// invalidate old refresh token
	s.refreshTokenRepo.InvalidateManyByUserUUID(user.UUID)

	// create refresh token
	refreshTokenExpiredAt := helper.TimeNowUTC().Add(time.Minute * time.Duration(config.Envs.REFRESH_TOKEN_EXP_MINS))
	newRefreshTokenObj := &model.RefreshToken{
		UUID:      uuid.New().String(),
		Token:     uuid.New().String(),
		UserID:    user.ID,
		UserUUID:  user.UUID,
		UsedAt:    nil,
		ExpiredAt: &refreshTokenExpiredAt,
	}
	logger.Debugf("new refresh token: %+v", newRefreshTokenObj)
	newRefreshTokenObj, err = s.refreshTokenRepo.Create(newRefreshTokenObj)
	if err != nil {
		logger.Errorf("error creating refresh token: %v", err)
		return nil, err
	}

	resp := &dto.RegisterUserRespData{
		AccessToken:  token,
		RefreshToken: newRefreshTokenObj.Token,
	}
	return resp, nil
}

func (s *AuthUcase) Login(payload dto.LoginReq) (*dto.LoginRespData, error) {
	// validate input
	err := payload.Validate()
	if err != nil {
		return nil, &error_utils.CustomErr{
			HttpCode: 400,
			Message:  err.Error(),
		}
	}

	// check if user exists
	var existing_user *model.User
	if strings.Contains(payload.UsernameOrEmail, "@") {
		existing_user, _ = s.userRepo.GetByEmail(payload.UsernameOrEmail)
	} else {
		existing_user, _ = s.userRepo.GetByUsername(payload.UsernameOrEmail)
	}
	if existing_user == nil {
		logger.Errorf("user not found")
		return nil, &error_utils.CustomErr{
			HttpCode: 401,
			Message:  "Invalid Credentials",
		}
	}
	logger.Debugf("user by username or email: %v", helper.PrettyJson(existing_user))

	// check password
	if !bcrypt_util.Compare(payload.Password, existing_user.Password) {
		logger.Errorf("invalid password")
		return nil, &error_utils.CustomErr{
			HttpCode: 401,
			Message:  "Invalid Credentials",
		}
	}

	// generate token
	token, err := jwt_util.GenerateJwtToken(existing_user, config.Envs.JWT_SECRET_KEY, config.Envs.JWT_EXP_MINS, nil)
	if err != nil {
		logger.Errorf("error generating token: %v", err)
		return nil, err
	}

	// invalidate old refresh token
	err = s.refreshTokenRepo.InvalidateManyByUserUUID(existing_user.UUID)
	if err != nil {
		logger.Errorf("error invalidating old refresh token: %v", err)
		return nil, err
	}

	// create refresh token
	refreshTokenExpiredAt := helper.TimeNowUTC().Add(time.Minute * time.Duration(config.Envs.REFRESH_TOKEN_EXP_MINS))
	newRefreshTokenObj := &model.RefreshToken{
		UUID:      uuid.New().String(),
		Token:     uuid.New().String(),
		UserID:    existing_user.ID,
		UserUUID:  existing_user.UUID,
		UsedAt:    nil,
		ExpiredAt: &refreshTokenExpiredAt,
	}
	logger.Debugf("new refresh token: %+v", helper.PrettyJson(newRefreshTokenObj))
	newRefreshTokenObj, err = s.refreshTokenRepo.Create(newRefreshTokenObj)
	if err != nil {
		logger.Errorf("error creating refresh token: %v", err)
		return nil, err
	}

	return &dto.LoginRespData{
		AccessToken:  token,
		RefreshToken: newRefreshTokenObj.Token,
	}, nil
}

func (s *AuthUcase) RefreshToken(payload dto.RefreshTokenReq) (*dto.RefreshTokenRespData, error) {
	// get refresh token
	refreshToken, err := s.refreshTokenRepo.GetByToken(payload.RefreshToken)
	logger.Debugf("refresh token: %+v", helper.PrettyJson(refreshToken))
	logger.Debugf("error: %v", err)
	if err != nil {
		logger.Errorf("refresh token not found: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 401,
			Message:  "Invalid Refresh Token",
		}
	}

	// check if refresh token is expired
	if refreshToken.ExpiredAt != nil {
		if refreshToken.ExpiredAt.Before(helper.TimeNowUTC()) {
			logger.Errorf("refresh token is expired")
			return nil, &error_utils.CustomErr{
				HttpCode: 401,
				Message:  "Invalid Refresh Token",
			}
		}
	}

	// check if refresh token is used
	if refreshToken.UsedAt != nil {
		logger.Errorf("refresh token is used")
		return nil, &error_utils.CustomErr{
			HttpCode: 401,
			Message:  "Invalid Refresh Token",
		}
	}

	// check if refresh token is valid
	if refreshToken.Invalid {
		logger.Errorf("refresh token is invalid")
		return nil, &error_utils.CustomErr{
			HttpCode: 401,
			Message:  "Invalid Refresh Token",
		}
	}

	// mark refresh token as used
	timeNow := helper.TimeNowUTC()
	refreshToken.UsedAt = &timeNow
	refreshToken, err = s.refreshTokenRepo.Update(refreshToken)
	if err != nil {
		logger.Errorf("error updating refresh token: %v", err)
		return nil, err
	}

	// get user
	user, err := s.userRepo.GetByID(refreshToken.UserID)
	if err != nil {
		logger.Errorf("user not found: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 500,
			Message:  "Internal server error",
			Detail:   err.Error(),
		}
	}

	// generate token
	token, err := jwt_util.GenerateJwtToken(user, config.Envs.JWT_SECRET_KEY, config.Envs.JWT_EXP_MINS, nil)
	if err != nil {
		logger.Errorf("error generating token: %v", err)
		return nil, err
	}

	// invalidate old refresh token
	err = s.refreshTokenRepo.InvalidateManyByUserUUID(user.UUID)
	if err != nil {
		logger.Errorf("error invalidating old refresh token: %v", err)
		return nil, err
	}

	// create refresh token
	refreshTokenExpiredAt := helper.TimeNowUTC().Add(time.Minute * time.Duration(config.Envs.REFRESH_TOKEN_EXP_MINS))
	newRefreshTokenObj := &model.RefreshToken{
		UUID:      uuid.New().String(),
		Token:     uuid.New().String(),
		UserID:    user.ID,
		UserUUID:  user.UUID,
		UsedAt:    nil,
		ExpiredAt: &refreshTokenExpiredAt,
	}
	newRefreshTokenObj, err = s.refreshTokenRepo.Create(newRefreshTokenObj)
	if err != nil {
		logger.Errorf("error creating refresh token: %v", err)
		return nil, err
	}

	return &dto.RefreshTokenRespData{
		AccessToken:  token,
		RefreshToken: newRefreshTokenObj.Token,
	}, nil
}

func (s *AuthUcase) CheckToken(payload dto.CheckTokenReq) (*dto.CheckTokenRespData, error) {
	claims, err := jwt_util.ValidateJWT(payload.AccessToken, config.Envs.JWT_SECRET_KEY)
	if err != nil || claims == nil {
		logger.Errorf("error validating token: %v", err)
		return nil, &error_utils.CustomErr{
			HttpCode: 401,
			Message:  "Invalid Access Token",
			Detail:   err.Error(),
		}
	}

	resp := &dto.CheckTokenRespData{
		UUID:     claims.UUID,
		Username: claims.Username,
		Role:     claims.Role,
		Email:    claims.Email,
	}

	return resp, nil
}
