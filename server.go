package main

import (
	"backend/domain/dto"
	"backend/handler"
	"backend/middleware"
	"backend/utils/http_response"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer(ginEngine *gin.Engine, deps CommonDeps) {
	responseWriter := http_response.NewHttpResponseWriter()

	// handlers
	authHandler := handler.NewAuthHandler(responseWriter, deps.AuthUcase)
	userHandler := handler.NewUserHandler(responseWriter, deps.UserUcase)
	loanHandler := handler.NewLoanHandler(responseWriter, deps.LoanUcase)

	// register routes
	router := ginEngine
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, dto.BaseJSONResp{
			Code:    200,
			Message: "pong",
		})
	})
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authHandler.Register)
		authRouter.POST("/login/dev", authHandler.LoginDev)
		authRouter.POST("/login", authHandler.Login)
		authRouter.POST("/check-token", authHandler.CheckToken)
		authRouter.POST("/refresh-token", authHandler.RefreshToken)
	}

	authMiddleware := middleware.AuthMiddleware(responseWriter)
	adminOnlyMiddleware := middleware.AuthAdminOnlyMiddleware(responseWriter)
	securedRouter := router.Group("/")
	securedRouter.Use(authMiddleware)
	{
		userRouter := securedRouter.Group("/users")
		{
			userRouter.GET("", userHandler.GetUserList)
			userRouter.GET("/me", userHandler.GetMe)
			userRouter.GET("/:uuid", userHandler.GetByUUID)
			userRouter.POST("", adminOnlyMiddleware, userHandler.CreateUser)
			userRouter.PUT("", userHandler.UpdateUserMe)
			userRouter.PUT("/:uuid", adminOnlyMiddleware, userHandler.UpdateUser)
			userRouter.DELETE("/:uuid", adminOnlyMiddleware, userHandler.DeleteUser)
			userRouter.POST("/ktp-photo", userHandler.UploadKtpPhoto)
			userRouter.POST("/face-photo", userHandler.UploadFacePhoto)
			userRouter.POST("/:uuid/current-limit", adminOnlyMiddleware, userHandler.UpdateCurrentLimit)
		}

		loanRouter := securedRouter.Group("/loans")
		{
			loanRouter.POST("", loanHandler.CreateNewLoan)
			loanRouter.POST("/:uuid/status", adminOnlyMiddleware, loanHandler.UpdateLoanStatus)
			loanRouter.GET("", loanHandler.GetLoanList)
		}
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/swagger/index.html")
	})
}
