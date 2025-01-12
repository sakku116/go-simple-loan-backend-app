package handler

import (
	"backend/domain/dto"
	ucase "backend/usecase"
	"backend/utils/http_response"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	respWriter http_response.IHttpResponseWriter
	loanUcase  ucase.ILoanUcase
}

type ILoanHandler interface {
	CreateNewLoan(c *gin.Context)
}

func NewLoanHandler(
	respWriter http_response.IHttpResponseWriter,
	loanUcase ucase.ILoanUcase,
) ILoanHandler {
	return LoanHandler{
		respWriter: respWriter,
		loanUcase:  loanUcase,
	}
}

// @Summary Create new loan (current user)
// @Tags Loan
// @Accept json
// @Produce json
// @Param loan body dto.CreateNewLoanReq true "Create loan rquest"
// @Success 200 {object} dto.BaseJSONResp{data=dto.CreateNewLoanRespData}
// @Security OAuth2Password
// @Router /loans [post]
func (h LoanHandler) CreateNewLoan(c *gin.Context) {
	var payload dto.CreateNewLoanReq
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

	data, err := h.loanUcase.CreateNewLoan(currentUser.UUID, payload)
	if err != nil {
		h.respWriter.HTTPCustomErr(c, err)
		return
	}

	h.respWriter.HTTPJsonOK(c, data)
}
