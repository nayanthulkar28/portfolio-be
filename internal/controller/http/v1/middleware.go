package v1

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	APIKey string
}

func NewMiddleware(apiKey string) *Middleware {
	return &Middleware{
		APIKey: apiKey,
	}
}

func (m *Middleware) ValidateApiKey(c *gin.Context) {
	apiKey := c.Request.Header.Get("p-api-key")
	decodeApiKey, err := base64.StdEncoding.DecodeString(apiKey)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid api key")
		return
	}
	if string(decodeApiKey) != m.APIKey {
		errorResponse(c, http.StatusBadRequest, "invalid api key")
		return
	}
	c.Next()
}
