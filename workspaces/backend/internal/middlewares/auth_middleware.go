package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(global.BearerTokenKey)
		bearerToken := strings.Split(token, " ")

		if len(bearerToken) != 2 || len(bearerToken[1]) == 0 {
			response.MessageResponse(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		claims, errCode := utils.VerifyJWT(bearerToken[1])

		if errCode != nil {
			response.MessageResponse(c, *errCode)
			c.Abort()
			return
		}

		c.Request.Header.Set("userId", claims.UserID)
		c.Request.Header.Set("userEmail", claims.Data["email"].(string))

		c.Next()
	}
}
