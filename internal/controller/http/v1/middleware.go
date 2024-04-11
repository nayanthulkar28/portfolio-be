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

func (m *Middleware) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Accept,X-Requested-With,Content-Type,p-api-key")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
