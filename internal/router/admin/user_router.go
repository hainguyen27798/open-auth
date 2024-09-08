package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/middlewares"
	"github.com/open-auth/internal/wires"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wires.InitUserRouterHandler()

	userRouterPrivate := Router.Group("/users")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		//userRouterPrivate.GET("")
		userRouterPrivate.GET("/me", userController.GetMe)
	}
}
