package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/wires"
)

type AuthenticateRouter struct{}

func (authenticateRouter *AuthenticateRouter) InitAuthenticateRouter(Router *gin.RouterGroup) {
	authController, _ := wires.InitAuthRouterHandler()

	// public router
	authenticateRouterPublic := Router.Group("/auth")
	{
		authenticateRouterPublic.POST("/register", authController.Register)
		authenticateRouterPublic.POST("/login", authController.Login)
		authenticateRouterPublic.POST("/login-admin", authController.LoginAdmin)
		authenticateRouterPublic.POST("/refresh-token", authController.RefreshToken)
		authenticateRouterPublic.POST("/logout", authController.Logout)
	}
}
