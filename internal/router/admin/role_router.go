package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/middlewares"
	"github.com/open-auth/internal/wires"
)

type RoleRouter struct{}

func (rr *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	RoleController, _ := wires.InitRoleRouterHandler()

	roleRouter := Router.Group("roles")
	roleRouter.Use(middlewares.AuthMiddleware())
	{
		roleRouter.POST("", RoleController.Create)
		roleRouter.GET("", RoleController.GetAll)
		roleRouter.GET(":id", RoleController.Get)
		roleRouter.PATCH(":id", RoleController.Update)
		roleRouter.DELETE(":id", RoleController.Delete)
	}
}
