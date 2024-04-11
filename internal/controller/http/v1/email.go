package v1

import (
	"net/http"
	"portfolio-be/internal/domain"
	"portfolio-be/internal/usecase"

	"github.com/gin-gonic/gin"
)

type EmailRoutes struct {
	eu usecase.EmailUsecase
}

func newEmailRoutes(handler *gin.RouterGroup, eu *usecase.EmailUsecase) {
	r := &EmailRoutes{*eu}
	h := handler.Group("/email")
	{
		h.POST("", r.SendEmail)
	}
}

func (r *EmailRoutes) SendEmail(c *gin.Context) {
	var req domain.SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, INVALID_REQUEST_BODY_MESSAGE)
		return
	}

	response, err := r.eu.SendEmail(c, req)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response)
}
