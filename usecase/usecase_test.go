package ucase

import (
	"backend/config"
	"backend/mocks"
)

var (
	MockedUserRepo         *mocks.IUserRepo
	MockedRefreshTokenRepo *mocks.IRefreshTokenRepo
	MockedFileRepo         *mocks.IFileRepo
	TestAuthUcase          IAuthUcase
	TestUserUcase          IUserUcase
)

func SetupTest() {
	config.InitEnv("../.env")
	config.ConfigureLogger()

	// repositories
	MockedUserRepo = new(mocks.IUserRepo)
	MockedRefreshTokenRepo = new(mocks.IRefreshTokenRepo)

	// usecases
	TestAuthUcase = NewAuthUcase(MockedUserRepo, MockedRefreshTokenRepo)
	TestUserUcase = NewUserUcase(MockedUserRepo, MockedFileRepo)
}
