package main

import (
	"backend/config"
	"backend/domain/model"
	"backend/repository"
	ucase "backend/usecase"
	file_storage_util "backend/utils/file"
	"backend/utils/helper"
	minio_util "backend/utils/minio"
	seeder_util "backend/utils/seeder/user"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

func init() {
	logger.Debugf("User: %v", helper.PrettyJson(model.User{}))
	config.InitEnv("./.env")
	config.ConfigureLogger()
}

var logger = logging.MustGetLogger("main")

// @title Loan Backend API
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl /auth/login/dev
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme (add 'Bearer ' prefix).
func main() {
	logger.Debugf("Envs: %v", helper.PrettyJson(config.Envs))

	gormDB := config.NewMySqlDB()
	minioClient := config.NewMinioClient()

	// migrations
	err := gormDB.AutoMigrate(
		&model.User{},
		&model.RefreshToken{},
		&model.Loan{},
	)
	if err != nil {
		logger.Fatalf("failed to migrate database: %v", err)
	}

	// ensure buckets
	err = minio_util.EnsureBucketsFromModels(minioClient, &model.User{})
	if err != nil {
		logger.Fatalf("failed to ensure buckets: %v", err)
	}

	// repositories
	userRepo := repository.NewUserRepo(gormDB)
	refreshTokenRepo := repository.NewRefreshTokenRepo(gormDB)
	loanRepo := repository.NewLoanRepo(gormDB)

	// utils
	fileStorageUtil := file_storage_util.NewFileStorageUtil(minioClient)

	// ucases
	authUcase := ucase.NewAuthUcase(userRepo, refreshTokenRepo)
	userUcase := ucase.NewUserUcase(userRepo, fileStorageUtil)
	loanUcase := ucase.NewLoanCase(loanRepo, userRepo)

	dependencies := CommonDeps{
		AuthUcase: authUcase,
		UserUcase: userUcase,
		LoanUcase: loanUcase,
	}

	// proccess args
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "--seed-user":
			logger.Infof("running seed user")
			err := seeder_util.SeedUser(userRepo)
			if err != nil {
				logger.Fatalf("failed to seed user: %v", err)
			}
			logger.Infof("seed user finished")
		}
	}

	ginEngine := gin.Default()
	SetupServer(ginEngine, dependencies)
	ginEngine.Run(fmt.Sprintf("%s:%d", config.Envs.HOST, config.Envs.PORT))
}
