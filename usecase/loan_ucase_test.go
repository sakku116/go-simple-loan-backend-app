package ucase

import (
	"backend/domain/dto"
	"backend/domain/enum"
	"backend/domain/model"
	error_utils "backend/utils/error"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestLoanUcase_CreateNewLoan(t *testing.T) {
	SetupTest()

	mockKtpPhoto := "test.jpg"
	mockFacePhoto := "test.jpg"
	t.Run("test success", func(t *testing.T) {
		userUUID := "test-uuid"
		payload := dto.CreateNewLoanReq{
			OTR:        5000,
			AssetName:  "Car",
			TermMonths: 2,
		}
		mockUser := &model.User{
			UUID:         userUUID,
			Model:        gorm.Model{ID: 1},
			CurrentLimit: 20000,
			KtpPhoto:     &mockKtpPhoto,
			FacePhoto:    &mockFacePhoto,
			Email:        "test@example.com",
			Username:     "Test User",
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(mockUser, nil).Once()
		MockedLoanRepo.On("GetUnPaidListByUserID", mockUser.ID).Return([]model.Loan{}, nil).Once()
		MockedLoanRepo.On("Create", mock.AnythingOfType("*model.Loan")).Return(&model.Loan{
			UUID:              "loan-uuid",
			UserID:            mockUser.ID,
			UserUUID:          mockUser.UUID,
			RefNumber:         123413515,
			OTR:               payload.OTR,
			AssetName:         payload.AssetName,
			TermMonths:        payload.TermMonths,
			Status:            enum.LoanStatus_PENDING,
			InterestRate:      500,
			AdminFee:          100,
			TotalAmount:       5700,
			InstallmentAmount: 475,
		}, nil).Once()

		resp, err := TestLoanUcase.CreateNewLoan(userUUID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, float64(14400), resp.CurrentLimitRemaining) // based on the calculation in the function
	})

	t.Run("test user not found", func(t *testing.T) {
		userUUID := "test-uuid"
		payload := dto.CreateNewLoanReq{
			OTR:        5000,
			AssetName:  "Car",
			TermMonths: 2,
		}

		MockedUserRepo.On("GetByUUID", userUUID).Return(nil, fmt.Errorf("not found")).Once()

		resp, err := TestLoanUcase.CreateNewLoan(userUUID, payload)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if customErr, ok := err.(*error_utils.CustomErr); ok {
			assert.Equal(t, 404, customErr.HttpCode)
			assert.Equal(t, "user not found", customErr.Message)
		}
	})

	t.Run("test insufficient limit", func(t *testing.T) {
		userUUID := "test-uuid"
		payload := dto.CreateNewLoanReq{
			OTR:        30000,
			AssetName:  "Car",
			TermMonths: 2,
		}
		mockUser := &model.User{
			UUID:         userUUID,
			Model:        gorm.Model{ID: 1},
			CurrentLimit: 20000,
			KtpPhoto:     &mockKtpPhoto,
			FacePhoto:    &mockFacePhoto,
			Email:        "test@example.com",
			Username:     "Test User",
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(mockUser, nil).Once()
		MockedLoanRepo.On("GetUnPaidListByUserID", mockUser.ID).Return([]model.Loan{}, nil).Once()

		resp, err := TestLoanUcase.CreateNewLoan(userUUID, payload)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if customErr, ok := err.(*error_utils.CustomErr); ok {
			assert.Equal(t, 400, customErr.HttpCode)
			assert.Equal(t, "limit not enough, current limit remaining: 20000", customErr.Message)
		}
	})

	t.Run("test missing photos", func(t *testing.T) {
		userUUID := "test-uuid"
		payload := dto.CreateNewLoanReq{
			OTR:        5000,
			AssetName:  "Car",
			TermMonths: 6,
		}
		mockUser := &model.User{
			UUID:         userUUID,
			Model:        gorm.Model{ID: 1},
			CurrentLimit: 20000,
		}
		MockedUserRepo.On("GetByUUID", userUUID).Return(mockUser, nil).Once()

		resp, err := TestLoanUcase.CreateNewLoan(userUUID, payload)
		assert.Error(t, err)
		assert.Nil(t, resp)
		if customErr, ok := err.(*error_utils.CustomErr); ok {
			assert.Equal(t, 400, customErr.HttpCode)
			assert.Equal(t, "Both KTP photo and face photo are required. Please update your account.", customErr.Message)
		}
	})
}
