package domain

const (
	EMPTY_EMAIL_NAME     = "ANONYMOUS"
	EMAIL_FAILED_STATUS  = "failed"
	EMAIL_SUCCESS_STATUS = "success"
)

type SendEmailRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type SendEmailResponse struct {
	Status string `json:"status"`
}
