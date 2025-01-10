package ucase

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"backend/domain/dto"
	"backend/domain/enum"
	"backend/domain/model"
	bcrypt_util "backend/utils/bcrypt"
	"backend/utils/helper"
)

func TestAuthUcase_Login(t *testing.T) {
	// setup
	SetupTest()

	password := "test-password"
	hashedPassword, _ := bcrypt_util.Hash(password)
	existingUser := &model.User{
		UUID:     "test-uuid",
		Email:    "test@gmail.com",
		Username: "test",
		Password: hashedPassword,
		NIK:      "1234567890123456",
	}

	// mock user repo
	MockedUserRepo.On("GetByEmail", existingUser.Email).Return(existingUser, nil).Once()
	MockedUserRepo.On("GetByUsername", existingUser.Username).Return(existingUser, nil)

	// mock refresh token repo
	MockedRefreshTokenRepo.On("InvalidateManyByUserUUID", existingUser.UUID).Return(nil)
	MockedRefreshTokenRepo.On("Create", mock.Anything).Return(
		&model.RefreshToken{},
		nil,
	)

	// test login by email
	t.Log("test login by email")
	payload := dto.LoginReq{
		UsernameOrEmail: existingUser.Email,
		Password:        password,
	}
	resp, err := TestAuthUcase.Login(payload)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// test login by username
	t.Log("test login by username")
	payload = dto.LoginReq{
		UsernameOrEmail: existingUser.Username,
		Password:        password,
	}
	resp, err = TestAuthUcase.Login(payload)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// test invalid credentials
	t.Log("test invalid credentials")
	MockedUserRepo.On("GetByEmail", "invalid@gmail.com").Return(nil, nil)
	payload = dto.LoginReq{
		UsernameOrEmail: "invalid@gmail.com",
		Password:        "wrong-password",
	}
	resp, err = TestAuthUcase.Login(payload)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "Invalid Credentials")
}

func TestAuthUcase_Register(t *testing.T) {
	// setup
	SetupTest()

	payload := dto.RegisterUserReq{
		Email:         "user1@gmail.com",
		Username:      "user1",
		Password:      "test-password",
		NIK:           "9087653421654321",
		Fullname:      "test1",
		Legalname:     "test1",
		Birthdate:     "11-06-2004",
		Birthplace:    "test1",
		CurrentSalary: 10000000,
	}

	// test success
	t.Log("test success")
	MockedUserRepo.On("GetByEmail", payload.Email).Return(nil, nil).Once()
	MockedUserRepo.On("GetByUsername", payload.Username).Return(nil, nil)
	MockedUserRepo.On("GetByNIK", payload.NIK).Return(nil, nil)
	newUserUUID := "test-uuid"
	hashedPassword, _ := bcrypt_util.Hash(payload.Password)
	newUser := &model.User{
		UUID:          newUserUUID,
		Email:         payload.Email,
		Username:      payload.Username,
		Password:      hashedPassword,
		NIK:           payload.NIK,
		Fullname:      payload.Fullname,
		Legalname:     payload.Legalname,
		Birthdate:     payload.Birthdate,
		Birthplace:    payload.Birthplace,
		CurrentSalary: payload.CurrentSalary,
		Role:          enum.UserRole_User,
	}
	MockedUserRepo.On("Create", mock.MatchedBy(func(user *model.User) bool {
		// ignore auto generated values
		return user.Username == payload.Username &&
			user.Email == payload.Email &&
			user.Role == enum.UserRole_User &&
			user.NIK == payload.NIK &&
			user.Fullname == payload.Fullname &&
			user.Legalname == payload.Legalname &&
			user.Birthdate == payload.Birthdate &&
			user.Birthplace == payload.Birthplace &&
			user.CurrentSalary == payload.CurrentSalary
	})).Return(newUser, nil)
	MockedRefreshTokenRepo.On("InvalidateManyByUserUUID", newUserUUID).Return(nil)
	MockedRefreshTokenRepo.On("Create", mock.Anything).Return(
		&model.RefreshToken{},
		nil,
	)
	resp, err := TestAuthUcase.Register(nil, payload)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// test user already exists
	MockedUserRepo.On("GetByEmail", payload.Email).Return(&model.User{}, nil).Once()
	t.Log("test user already exists")
	resp, err = TestAuthUcase.Register(nil, payload)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "already exists")

}

func TestAuthUcase_RefreshToken(t *testing.T) {
	// setup
	SetupTest()

	// data
	payload := dto.RefreshTokenReq{
		RefreshToken: "valid-token",
	}

	// test success
	t.Log("test success")
	expiredAt := helper.TimeNowUTC().Add(1 * time.Hour)
	existingRefreshToken := &model.RefreshToken{
		UUID:      "test-uuid",
		UserID:    1,
		UserUUID:  "test-user-uuid",
		Token:     "valid-token",
		UsedAt:    nil,
		ExpiredAt: &expiredAt,
		Invalid:   false,
	}
	MockedRefreshTokenRepo.On("GetByToken", payload.RefreshToken).Return(existingRefreshToken, nil).Once()
	MockedRefreshTokenRepo.On("Update", mock.MatchedBy(func(token *model.RefreshToken) bool {
		// ignore UsedAt
		return token.UUID == existingRefreshToken.UUID &&
			token.UserID == existingRefreshToken.UserID &&
			token.UserUUID == existingRefreshToken.UserUUID &&
			token.Token == existingRefreshToken.Token &&
			token.ExpiredAt == existingRefreshToken.ExpiredAt &&
			token.Invalid == existingRefreshToken.Invalid
	})).Return(existingRefreshToken, nil)
	refreshTokenUser := &model.User{
		UUID: existingRefreshToken.UserUUID,
	}
	MockedUserRepo.On("GetByID", existingRefreshToken.UserID).Return(refreshTokenUser, nil)
	MockedRefreshTokenRepo.On("InvalidateManyByUserUUID", existingRefreshToken.UserUUID).Return(nil)
	expiredAt = helper.TimeNowUTC().Add(1 * time.Hour)
	MockedRefreshTokenRepo.On("Create", mock.Anything).Return(&model.RefreshToken{}, nil)
	resp, err := TestAuthUcase.RefreshToken(payload)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// test invalid token
	t.Log("test invalid token")
	MockedRefreshTokenRepo.On("GetByToken", payload.RefreshToken).Return(nil, errors.New("not found")).Once()
	resp, err = TestAuthUcase.RefreshToken(payload)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "Invalid Refresh Token")
}
