package usecase

import (
	"portfolio-be/internal/domain"
	webapis "portfolio-be/internal/webAPIs"

	"github.com/gin-gonic/gin"
)

type EmailUsecase struct {
	ew webapis.EmailWebAPI
}

func NewEmailUsecase(ew webapis.EmailWebAPI) *EmailUsecase {
	return &EmailUsecase{
		ew: ew,
	}
}

func (eu *EmailUsecase) SendEmail(ctx *gin.Context, request domain.SendEmailRequest) (domain.SendEmailResponse, error) {
	var response domain.SendEmailResponse
	if request.Name == "" {
		request.Name = domain.EMPTY_EMAIL_NAME
	}

	err := eu.ew.SendEmail(ctx, request)
	if err != nil {
		response.Status = domain.EMAIL_FAILED_STATUS
		return response, err
	}

	response.Status = domain.EMAIL_SUCCESS_STATUS
	return response, nil
}
