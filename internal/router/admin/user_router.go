package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/middlewares"
	"github.com/open-auth/internal/wires"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wires.InitUserRouterHandler()

	userRouterMePrivate := Router.Group("/users")
	userRouterMePrivate.Use(middlewares.AuthMiddleware(global.UserScope, global.AdminScope))
	{
		userRouterMePrivate.GET("/me", userController.GetMe)
	}

	userRouterAdminPrivate := Router.Group("/users")
	userRouterAdminPrivate.Use(middlewares.AuthMiddleware(global.AdminScope))
	{
		userRouterAdminPrivate.GET("/", userController.GetUsers)
	}
}
