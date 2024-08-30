package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/pkg/response"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			response.ErrorResponse(c, response.ErrInvalidToken, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
