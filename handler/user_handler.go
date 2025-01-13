package handler

import (
	"backend/domain/dto"
	ucase "backend/usecase"
	"backend/utils/http_response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	respWriter http_response.IHttpResponseWriter
	userUcase  ucase.IUserUcase
}

type IUserHandler interface {
	GetByUUID(c *gin.Context)
	GetMe(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateUserMe(c *gin.Context)
	DeleteUser(c *gin.Context)
	UploadKtpPhoto(c *gin.Context)
	UploadFacePhoto(c *gin.Context)
	GetUserList(c *gin.Context)
	UpdateCurrentLimit(c *gin.Context)
}

func NewUserHandler(
	respWriter http_response.IHttpResponseWriter,
	userUcase ucase.IUserUcase,
) IUserHandler {
	return UserHandler{
		respWriter: respWriter,
		userUcase:  userUcase,
	}
}

// @Summary Get user by UUID (admin only)
// @Tags User
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} dto.BaseJSONResp{data=dto.GetUserByUUIDResp}
// @Security OAuth2Password
// @Router /users/{uuid} [get]
func (h UserHandler) GetByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.userUcase.GetByUUID(c, uuid)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Get user (current user)
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} dto.BaseJSONResp{data=dto.GetUserByUUIDResp}
// @Security OAuth2Password
// @Router /users/me [get]
func (h UserHandler) GetMe(c *gin.Context) {
	// get current user
	tmp, ok := c.Get("currentUser")
	if !ok {
		logger.Error("currentUser not found")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser not found", nil)
		return
	}
	currentUser, ok := tmp.(*dto.CurrentUser)
	if !ok {
		logger.Error("currentUser type missmatched")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser type missmatched", nil)
		return
	}

	resp, err := h.userUcase.GetByUUID(c, currentUser.UUID)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Create user (admin only)
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.CreateUserReq true "User create request"
// @Success 200 {object} dto.BaseJSONResp{data=dto.CreateUserRespData}
// @Security OAuth2Password
// @Router /users [post]
func (h UserHandler) CreateUser(c *gin.Context) {
	var payload dto.CreateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}

	resp, err := h.userUcase.CreateUser(c, payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Update user by UUID (admin only)
// @Tags User
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Param user body dto.UpdateUserReq true "User update request"
// @Success 200 {object} dto.BaseJSONResp{data=dto.UpdateUserRespData}
// @Security OAuth2Password
// @Router /users/{uuid} [put]
func (h UserHandler) UpdateUser(c *gin.Context) {
	uuid := c.Param("uuid")
	var payload dto.UpdateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}

	resp, err := h.userUcase.UpdateUser(c, uuid, payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Update user (current user)
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UpdateUserReq true "User update request"
// @Success 200 {object} dto.BaseJSONResp{data=dto.UpdateUserRespData}
// @Security OAuth2Password
// @Router /users [put]
func (h UserHandler) UpdateUserMe(c *gin.Context) {
	var payload dto.UpdateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}

	// get current user
	tmp, ok := c.Get("currentUser")
	if !ok {
		logger.Error("currentUser not found")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser not found", nil)
		return
	}
	currentUser, ok := tmp.(*dto.CurrentUser)
	if !ok {
		logger.Error("currentUser type missmatched")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser type missmatched", nil)
		return
	}

	resp, err := h.userUcase.UpdateUser(c, currentUser.UUID, payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Delete user by UUID (admin only)
// @Tags User
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} dto.BaseJSONResp{data=dto.DeleteUserRespData}
// @Security OAuth2Password
// @Router /users/{uuid} [delete]
func (h UserHandler) DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.userUcase.DeleteUser(c, uuid)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Upload my KTP photo (current user)
// @Tags User
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "KTP photo file"
// @Success 200 {object} dto.BaseJSONResp{data=dto.UploadKtpPhotoRespData}
// @Security OAuth2Password
// @Router /users/ktp-photo [post]
func (h UserHandler) UploadKtpPhoto(c *gin.Context) {
	var payload dto.UploadKtpPhotoReq
	if err := c.ShouldBind(&payload); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}

	// get current user
	tmp, ok := c.Get("currentUser")
	if !ok {
		logger.Error("currentUser not found")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser not found", nil)
		return
	}
	currentUser, ok := tmp.(*dto.CurrentUser)
	if !ok {
		logger.Error("currentUser type missmatched")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser type missmatched", nil)
		return
	}

	resp, err := h.userUcase.UploadKtpPhoto(c, currentUser.UUID, payload.File)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Upload my Face photo (current user)
// @Tags User
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Face photo file"
// @Success 200 {object} dto.BaseJSONResp{data=dto.UploadFacePhotoRespData}
// @Security OAuth2Password
// @Router /users/face-photo [post]
func (h UserHandler) UploadFacePhoto(c *gin.Context) {
	var payload dto.UploadFacePhotoReq
	if err := c.ShouldBind(&payload); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}

	// get current user
	tmp, ok := c.Get("currentUser")
	if !ok {
		logger.Error("currentUser not found")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser not found", nil)
		return
	}
	currentUser, ok := tmp.(*dto.CurrentUser)
	if !ok {
		logger.Error("currentUser type missmatched")
		h.respWriter.HTTPJson(c, 500, "internal server error", "currentUser type missmatched", nil)
		return
	}

	resp, err := h.userUcase.UploadFacePhoto(c, currentUser.UUID, payload.File)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Get user list (admin only)
// @Description Get user list
// @Tags User
// @Accept json
// @Produce json
// @Param params query dto.GetUserListReq true "Query parameters"
// @Success 200 {object} dto.BaseJSONResp{data=dto.GetUserListRespData}
// @Security OAuth2Password
// @Router /users [get]
func (h UserHandler) GetUserList(c *gin.Context) {
	var params dto.GetUserListReq
	if err := c.ShouldBindQuery(&params); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}
	resp, err := h.userUcase.GetUserList(c, params)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}
	h.respWriter.HTTPJsonOK(c, resp)
}

// @Summary Update user current limit (admin only)
// @Tags User
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Param payload body dto.UpdateCurrentLimitReq true "payload"
// @Success 200 {object} dto.BaseJSONResp{data=dto.UpdateCurrentLimitRespData}
// @Security OAuth2Password
// @Router /users/{uuid}/current-limit [post]
func (h UserHandler) UpdateCurrentLimit(c *gin.Context) {
	userUUID := c.Param("uuid")

	var payload dto.UpdateCurrentLimitReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		h.respWriter.HTTPJson(c, 400, "invalid payload", err.Error(), nil)
		return
	}

	resp, err := h.userUcase.UpdateCurrentLimit(c, userUUID, payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}

	h.respWriter.HTTPJsonOK(c, resp)
}
