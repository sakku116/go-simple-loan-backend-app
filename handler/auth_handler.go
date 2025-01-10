package handler

import (
	"backend/domain/dto"
	ucase "backend/usecase"
	"backend/utils/http_response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	respWriter http_response.IHttpResponseWriter
	authUcase  ucase.IAuthUcase
}

type IAuthHandler interface {
}

func NewAuthHandler(respWriter http_response.IHttpResponseWriter, authUcase ucase.IAuthUcase) AuthHandler {
	return AuthHandler{
		respWriter: respWriter,
		authUcase:  authUcase,
	}
}

// Register
// @Summary register new user
// @Tags Auth
// @Success 200 {object} dto.BaseJSONResp{data=dto.RegisterUserRespData}
// @Router /auth/register [post]
// @param payload  body  dto.RegisterUserReq  true "payload"
func (h *AuthHandler) Register(ctx *gin.Context) {
	var payload dto.RegisterUserReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		h.respWriter.HTTPJson(ctx, 400, "invalid payload", err.Error(), nil)
		return
	}

	data, err := h.authUcase.Register(ctx, payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	h.respWriter.HTTPJsonOK(ctx, data)
}

// Login Dev
// @Summary login dev
// @Tags Auth
// @Success 200 {object} dto.BaseJSONResp{data=dto.LoginDevResp{data=dto.LoginRespData}}
// @Router /auth/login/dev [post]
// @param payload  body  dto.LoginDevReq  true "payload"
func (h *AuthHandler) LoginDev(ctx *gin.Context) {
	var payload dto.LoginDevReq
	err := ctx.ShouldBind(&payload)
	if err != nil {
		h.respWriter.HTTPJson(ctx, 400, "invalid payload", err.Error(), nil)
		return
	}

	// convert payload
	newPayload := dto.LoginReq{
		UsernameOrEmail: payload.Username,
		Password:        payload.Password,
	}

	data, err := h.authUcase.Login(newPayload)
	if err != nil {
		h.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	// convert resp
	resp := dto.LoginDevResp{
		BaseJSONResp: dto.BaseJSONResp{Code: 200, Message: "success", Detail: "", Data: data},
		AccessToken:  data.AccessToken,
	}

	ctx.JSON(200, resp)
}

// Login
// @Summary login
// @Tags Auth
// @Success 200 {object} dto.BaseJSONResp{data=dto.LoginRespData}
// @Router /auth/login [post]
// @param payload  body  dto.LoginReq  true "payload"
func (h *AuthHandler) Login(ctx *gin.Context) {
	var payload dto.LoginReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		h.respWriter.HTTPJson(ctx, 400, "invalid payload", err.Error(), nil)
		return
	}

	data, err := h.authUcase.Login(payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	h.respWriter.HTTPJsonOK(ctx, data)
}

// Refresh Token
// @Tags Auth
// @Router /auth/refresh-token [post]
// @param payload  body  dto.RefreshTokenReq  true "payload"
// @Success 200 {object} dto.BaseJSONResp{data=dto.RefreshTokenRespData}
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	var payload dto.RefreshTokenReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		h.respWriter.HTTPJson(ctx, 400, "invalid payload", err.Error(), nil)
		return
	}

	data, err := h.authUcase.RefreshToken(payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	h.respWriter.HTTPJsonOK(ctx, data)
}

// Check Token
// @Tags Auth
// @Router /auth/check-token [post]
// @param payload  body  dto.CheckTokenReq  true "payload"
// @Success 200 {object} dto.BaseJSONResp{data=dto.CheckTokenRespData}
func (h AuthHandler) CheckToken(ctx *gin.Context) {
	var payload dto.CheckTokenReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		h.respWriter.HTTPJson(ctx, 400, "invalid payload", err.Error(), nil)
		return
	}

	data, err := h.authUcase.CheckToken(payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	h.respWriter.HTTPJsonOK(ctx, data)
}
