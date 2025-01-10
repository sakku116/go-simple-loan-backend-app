package main

import (
	"backend/domain/dto"
	"backend/handler"
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
	_ = authHandler

	// register routes
	router := ginEngine
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, dto.BaseJSONResp{
			Code:    200,
			Message: "pong",
		})
	})
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login/dev", authHandler.LoginDev)
	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/check-token", authHandler.CheckToken)
	router.POST("/auth/refresh-token", authHandler.RefreshToken)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/swagger/index.html")
	})
}
