package ucase

import (
	"backend/config"
	"backend/mocks"
)

var (
	MockedFileStorageUtil  *mocks.IFileStorageUtil
	MockedUserRepo         *mocks.IUserRepo
	MockedRefreshTokenRepo *mocks.IRefreshTokenRepo
	MockedLoanRepo         *mocks.ILoanRepo
	TestAuthUcase          IAuthUcase
	TestUserUcase          IUserUcase
	TestLoanUcase          ILoanUcase
)

func SetupTest() {
	config.InitEnv("../.env")
	config.ConfigureLogger()
	MockedFileStorageUtil = new(mocks.IFileStorageUtil)

	// repositories
	MockedUserRepo = new(mocks.IUserRepo)
	MockedRefreshTokenRepo = new(mocks.IRefreshTokenRepo)
	MockedLoanRepo = new(mocks.ILoanRepo)

	// usecases
	TestAuthUcase = NewAuthUcase(MockedUserRepo, MockedRefreshTokenRepo)
	TestUserUcase = NewUserUcase(MockedUserRepo, MockedFileStorageUtil)
	TestLoanUcase = NewLoanCase(MockedLoanRepo, MockedUserRepo)
}
