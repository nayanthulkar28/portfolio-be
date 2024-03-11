package app

import (
	"portfolio-be/config"
	v1 "portfolio-be/internal/controller/http/v1"
	"portfolio-be/internal/usecase"
	webapis "portfolio-be/internal/webAPIs"

	"github.com/gin-gonic/gin"
)

// Run is use to setup your database connection, servers, usecase etc...
func Run(cfg *config.Config) {
	router := gin.Default()
	middleware := v1.NewMiddleware(cfg.App.APIKey)
	emailWebApi := webapis.NewEmailWebAPI(cfg.WebAPI.Email.Name, cfg.WebAPI.Email.FromEmailAddress, cfg.WebAPI.Email.FromEmailPassword, cfg.WebAPI.Email.ToEmailAddress, cfg.WebAPI.Email.SmtpAuthAddress, cfg.WebAPI.Email.SmtpServerAddress)
	emailUsecase := usecase.NewEmailUsecase(*emailWebApi)
	v1.NewRouter(router, middleware, emailUsecase)
	router.Run()
}
