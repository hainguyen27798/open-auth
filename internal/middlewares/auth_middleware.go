package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/global"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
	"golang.org/x/exp/slices"
	"strings"
)

func AuthMiddleware(scopes ...global.Scope) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(global.BearerTokenKey)
		bearerToken := strings.Split(token, " ")

		if len(bearerToken) != 2 || len(bearerToken[1]) == 0 {
			response.MessageResponse(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		tokenScope, err := utils.GetValueFromToken(bearerToken[1], "scope")
		currentScope := global.Scope(*tokenScope)
		if err != nil {
			response.MessageResponse(c, response.ErrInvalidToken)
			c.Abort()
			return
		}

		if slices.Contains(scopes, currentScope) == false {
			response.MessageResponse(c, response.ErrUnauthorized)
			c.Abort()
			return
		}

		claims, errCode := utils.VerifyJWT(currentScope, bearerToken[1])

		if errCode != nil {
			response.MessageResponse(c, errCode.Code())
			c.Abort()
			return
		}

		c.Request.Header.Set("userId", claims.UserID)
		c.Request.Header.Set("userEmail", claims.Data["email"].(string))

		c.Next()
	}
}
