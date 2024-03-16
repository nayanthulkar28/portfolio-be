package v1

import (
	"net/http"
	"portfolio-be/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, m *Middleware, eu *usecase.EmailUsecase) {
	h := handler.Group("portfolio-be/v1", m.ValidateApiKey)
	h.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "service is up!!!")
	})
	{
		newEmailRoutes(h, eu)
	}
}
