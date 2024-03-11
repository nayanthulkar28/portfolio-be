package webapis

import (
	"fmt"
	"net/smtp"
	"portfolio-be/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

type EmailWebAPI struct {
	Name              string
	FromEmailAddress  string
	FromEmailPassword string
	ToEmailAddress    string
	SmtpAuthAddress   string
	SmtpServerAddress string
}

func NewEmailWebAPI(name string, fromEmailAddress string, fromEmailPassword string, toEmailAddress string, smtpAuth string, smtpServer string) *EmailWebAPI {
	return &EmailWebAPI{
		Name:              name,
		FromEmailAddress:  fromEmailAddress,
		FromEmailPassword: fromEmailPassword,
		ToEmailAddress:    toEmailAddress,
		SmtpAuthAddress:   smtpAuth,
		SmtpServerAddress: smtpServer,
	}
}

func (ew *EmailWebAPI) SendEmail(ctx *gin.Context, request domain.SendEmailRequest) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", ew.Name, ew.FromEmailAddress)
	e.To = []string{ew.ToEmailAddress}
	e.Subject = fmt.Sprintf("Mail from %s <%s>", request.Name, request.Email)
	e.Text = []byte(request.Message)

	smtpAuth := smtp.PlainAuth("", ew.FromEmailAddress, ew.FromEmailPassword, ew.SmtpAuthAddress)
	return e.Send(ew.SmtpServerAddress, smtpAuth)
}
