package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/wires"
)

type AuthenticateRouter struct{}

func (authenticateRouter *AuthenticateRouter) InitAuthenticateRouter(Router *gin.RouterGroup) {
	userController, _ := wires.InitUserRouterHandler()

	// public router
	authenticateRouterPublic := Router.Group("/auth")
	{
		authenticateRouterPublic.POST("/register", userController.Register)
		authenticateRouterPublic.POST("/login", userController.Login)
	}
}
