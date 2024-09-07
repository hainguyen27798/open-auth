package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/middlewares"
	"github.com/go-open-auth/internal/wires"
)

type RoleRouter struct{}

func (rr *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	RoleController, _ := wires.InitRoleRouterHandler()

	roleRouter := Router.Group("roles")
	roleRouter.Use(middlewares.AuthMiddleware())
	{
		roleRouter.POST("", RoleController.Create)
	}
}
