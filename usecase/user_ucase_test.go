package ucase

import (
	"backend/domain/dto"
	"backend/domain/enum"
	"backend/domain/model"
	bcrypt_util "backend/utils/bcrypt"
	error_utils "backend/utils/error"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUcase_GetByUUID(t *testing.T) {
	// setup
	SetupTest()

	// test success
	t.Log("test success")
	userUUID := "test-uuid"
	user := &model.User{
		UUID: userUUID,
	}
	MockedUserRepo.On("GetByUUID", userUUID).Return(user, nil).Once()
	resp, err := TestUserUcase.GetByUUID(context.Background(), userUUID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// test not found
	t.Log("test not found")
	MockedUserRepo.On("GetByUUID", userUUID).Return(nil, errors.New("not found")).Once()
	resp, err = TestUserUcase.GetByUUID(context.Background(), userUUID)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUserUcase_CreateUser(t *testing.T) {
	SetupTest()

	// test success
	t.Log("test success")
	payload := dto.CreateUserReq{
		Username:      "user1",
		Email:         "user1@gmail.com",
		Password:      "password",
		Role:          enum.UserRole_User,
		Fullname:      "User 1",
		Legalname:     "User 1",
		NIK:           "1234567890123456",
		Birthplace:    "Jakarta",
		Birthdate:     "02-06-2000",
		CurrentSalary: 12000000,
	}

	MockedUserRepo.On("GetByEmail", payload.Email).Return(nil, nil).Once()
	MockedUserRepo.On("GetByUsername", payload.Username).Return(nil, nil).Once()
	MockedUserRepo.On("GetByNIK", payload.NIK).Return(nil, nil).Once()
	hashedPassword, _ := bcrypt_util.Hash(payload.Password)
	newUser := &model.User{
		UUID:          "test-uuid",
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
		CurrentLimit:  payload.CurrentSalary / 3,
	}
	MockedUserRepo.On("Create", mock.AnythingOfType("*model.User")).Return(newUser, nil).Once()

	resp, err := TestUserUcase.CreateUser(context.Background(), payload)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	// test invalid email
	t.Log("test invalid email")
	payload = dto.CreateUserReq{
		Email: "invalid-email",
	}
	resp, err = TestUserUcase.CreateUser(context.Background(), payload)
	assert.Error(t, err)
	assert.Nil(t, resp)

	// test existing email
	t.Log("test existing email")
	payload = dto.CreateUserReq{
		Email: "user1@gmail.com",
	}
	MockedUserRepo.On("GetByEmail", payload.Email).Return(newUser, nil).Once()
	resp, err = TestUserUcase.CreateUser(context.Background(), payload)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUserUcase_UpdateUser(t *testing.T) {
	SetupTest()

	t.Run("test success", func(t *testing.T) {
		fullname := "User 2"
		legalname := "User 2"
		payload := dto.UpdateUserReq{
			Fullname:  &fullname,
			Legalname: &legalname,
		}
		userUUID := "test-uuid"
		existingUser := &model.User{
			UUID:      userUUID,
			Fullname:  "User 1",
			Legalname: "User 1",
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil)
		MockedUserRepo.On("GetByEmail", payload.Email).Return(nil, nil).Once()
		MockedUserRepo.On("GetByUsername", payload.Username).Return(nil, nil).Once()
		MockedUserRepo.On("GetByNIK", payload.NIK).Return(nil, nil).Once()

		existingUser.Legalname = *payload.Legalname
		existingUser.Fullname = *payload.Fullname
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(existingUser, nil)

		resp, err := TestUserUcase.UpdateUser(context.Background(), userUUID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("test user not found", func(t *testing.T) {
		payload := dto.UpdateUserReq{}
		MockedUserRepo.On("GetByUUID", "unexistinguseruuid").Return(nil, errors.New("not found")).Once()
		resp, err := TestUserUcase.UpdateUser(context.Background(), "unexistinguseruuid", payload)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestUserUcase_UploadKtpPhoto(t *testing.T) {
	SetupTest()

	t.Run("test success", func(t *testing.T) {
		userUUID := "test-uuid"
		file := &multipart.FileHeader{
			Filename: "ktp-photo.jpg",
		}
		existingUser := &model.User{
			UUID: userUUID,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("Upload", mock.Anything, file, mock.Anything, mock.Anything).Return("uploaded-ktp-photo.jpg", nil).Once()
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("GetUrl", mock.Anything, "uploaded-ktp-photo.jpg", mock.Anything, false).Return("http://bucket-name/uploaded-ktp-photo.jpg", nil).Once()

		resp, err := TestUserUcase.UploadKtpPhoto(context.Background(), userUUID, file)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "http://bucket-name/uploaded-ktp-photo.jpg", resp.KtpPhoto)
	})

	t.Run("test user not found", func(t *testing.T) {
		userUUID := "nonexistent-uuid"
		file := &multipart.FileHeader{
			Filename: "ktp-photo.jpg",
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(nil, fmt.Errorf("not found")).Once()

		resp, err := TestUserUcase.UploadKtpPhoto(context.Background(), userUUID, file)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("test file upload failure", func(t *testing.T) {
		userUUID := "test-uuid"
		file := &multipart.FileHeader{
			Filename: "ktp-photo.jpg",
		}
		existingUser := &model.User{
			UUID: userUUID,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("Upload", mock.Anything, file, mock.Anything, mock.Anything).Return("", fmt.Errorf("upload failed")).Once()

		resp, err := TestUserUcase.UploadKtpPhoto(context.Background(), userUUID, file)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("test update user failure", func(t *testing.T) {
		userUUID := "test-uuid"
		file := &multipart.FileHeader{
			Filename: "ktp-photo.jpg",
		}
		existingUser := &model.User{
			UUID: userUUID,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("Upload", mock.Anything, file, mock.Anything, mock.Anything).Return("uploaded-ktp-photo.jpg", nil).Once()
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(nil, fmt.Errorf("update failed")).Once()

		resp, err := TestUserUcase.UploadKtpPhoto(context.Background(), userUUID, file)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestUserUcase_UploadFacePhoto(t *testing.T) {
	SetupTest()

	t.Run("test success", func(t *testing.T) {
		userUUID := "test-uuid"
		file := &multipart.FileHeader{
			Filename: "face-photo.jpg",
		}
		existingUser := &model.User{
			UUID: userUUID,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("Upload", mock.Anything, file, mock.Anything, mock.Anything).Return("uploaded-face-photo.jpg", nil).Once()
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("GetUrl", mock.Anything, "uploaded-face-photo.jpg", mock.Anything, false).Return("http://bucket-name/uploaded-face-photo.jpg", nil).Once()

		resp, err := TestUserUcase.UploadFacePhoto(context.Background(), userUUID, file)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "http://bucket-name/uploaded-face-photo.jpg", resp.FacePhoto)
	})

	t.Run("test user not found", func(t *testing.T) {
		userUUID := "nonexistent-uuid"
		file := &multipart.FileHeader{
			Filename: "face-photo.jpg",
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(nil, fmt.Errorf("not found")).Once()

		resp, err := TestUserUcase.UploadFacePhoto(context.Background(), userUUID, file)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("test file upload failure", func(t *testing.T) {
		userUUID := "test-uuid"
		file := &multipart.FileHeader{
			Filename: "face-photo.jpg",
		}
		existingUser := &model.User{
			UUID: userUUID,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("Upload", mock.Anything, file, mock.Anything, mock.Anything).Return("", fmt.Errorf("upload failed")).Once()

		resp, err := TestUserUcase.UploadFacePhoto(context.Background(), userUUID, file)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("test update user failure", func(t *testing.T) {
		userUUID := "test-uuid"
		file := &multipart.FileHeader{
			Filename: "face-photo.jpg",
		}
		existingUser := &model.User{
			UUID: userUUID,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(existingUser, nil).Once()
		MockedFileStorageUtil.On("Upload", mock.Anything, file, mock.Anything, mock.Anything).Return("uploaded-face-photo.jpg", nil).Once()
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(nil, fmt.Errorf("update failed")).Once()

		resp, err := TestUserUcase.UploadFacePhoto(context.Background(), userUUID, file)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestUserUcase_GetUserList(t *testing.T) {
	SetupTest()

	t.Run("test success", func(t *testing.T) {
		params := dto.GetUserListReq{
			Page:  1,
			Limit: 10,
		}

		mockUserList := []model.User{
			{UUID: "user-uuid-1", Email: "user1@example.com"},
			{UUID: "user-uuid-2", Email: "user2@example.com"},
		}
		mockTotalData := int64(2)

		MockedUserRepo.On("GetList", mock.AnythingOfType("dto.UserRepo_GetListParams")).Return(mockUserList, mockTotalData, nil).Once()

		resp, err := TestUserUcase.GetUserList(context.Background(), params)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 2, len(resp.Data))
		assert.Equal(t, mockTotalData, resp.Total)
	})

	t.Run("test internal server error", func(t *testing.T) {
		params := dto.GetUserListReq{
			Page:  1,
			Limit: 10,
		}

		MockedUserRepo.On("GetList", mock.AnythingOfType("dto.UserRepo_GetListParams")).Return(nil, int64(0), fmt.Errorf("database error"))

		resp, err := TestUserUcase.GetUserList(context.Background(), params)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if customErr, ok := err.(*error_utils.CustomErr); ok {
			assert.Equal(t, 500, customErr.HttpCode)
			assert.Equal(t, "internal server error", customErr.Message)
		}
	})
}

func TestUserUcase_UpdateCurrentLimit(t *testing.T) {
	SetupTest()

	t.Run("test success", func(t *testing.T) {
		userUUID := "test-uuid"
		mockUser := &model.User{
			UUID:         userUUID,
			CurrentLimit: 100,
		}
		payload := dto.UpdateCurrentLimitReq{
			CurrentLimit: 200,
		}

		MockedUserRepo.On("GetByUUID", userUUID).Return(mockUser, nil).Once()
		mockUser.CurrentLimit = payload.CurrentLimit
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(mockUser, nil).Once()

		resp, err := TestUserUcase.UpdateCurrentLimit(context.Background(), userUUID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, payload.CurrentLimit, resp.CurrentLimit)
	})

	t.Run("test user not found", func(t *testing.T) {
		userUUID := "test-uuid"
		payload := dto.UpdateCurrentLimitReq{
			CurrentLimit: 200,
		}

		MockedUserRepo.On("GetByUUID", userUUID).Return(nil, fmt.Errorf("not found")).Once()

		resp, err := TestUserUcase.UpdateCurrentLimit(context.Background(), userUUID, payload)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if customErr, ok := err.(*error_utils.CustomErr); ok {
			assert.Equal(t, 404, customErr.HttpCode)
			assert.Equal(t, "user not found", customErr.Message)
		}
	})

	t.Run("test internal server error on update", func(t *testing.T) {
		userUUID := "test-uuid"
		payload := dto.UpdateCurrentLimitReq{
			CurrentLimit: 200,
		}

		mockUser := &model.User{
			UUID:         userUUID,
			CurrentLimit: 100,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(mockUser, nil).Once()
		MockedUserRepo.On("Update", mock.AnythingOfType("*model.User")).Return(nil, fmt.Errorf("database error")).Once()

		resp, err := TestUserUcase.UpdateCurrentLimit(context.Background(), userUUID, payload)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if customErr, ok := err.(*error_utils.CustomErr); ok {
			assert.Equal(t, 500, customErr.HttpCode)
			assert.Equal(t, "internal server error", customErr.Message)
		}
	})
}
