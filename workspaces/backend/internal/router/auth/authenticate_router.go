package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/wires"
)

type AuthenticateRouter struct{}

func (authenticateRouter *AuthenticateRouter) InitAuthenticateRouter(Router *gin.RouterGroup) {
	authController, _ := wires.InitAuthRouterHandler()

	// public router
	authenticateRouterPublic := Router.Group("/auth")
	{
		authenticateRouterPublic.POST("/register", authController.Register)
		authenticateRouterPublic.POST("/login", authController.Login)
	}
}
