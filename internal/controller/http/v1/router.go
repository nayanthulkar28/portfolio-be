package v1

import (
	"portfolio-be/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, m *Middleware, eu *usecase.EmailUsecase) {
	h := handler.Group("portfolio-be/v1", m.ValidateApiKey)
	{
		newEmailRoutes(h, eu)
	}
}
