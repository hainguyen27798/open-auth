package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/middlewares"
	"github.com/open-auth/internal/wires"
)

type PermissionRouter struct{}

func (pr *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	permissionController, _ := wires.InitPermissionRouterHandler()

	permissionRouter := Router.Group("permissions")
	permissionRouter.Use(middlewares.AuthMiddleware(global.AdminScope))
	{
		permissionRouter.POST("", permissionController.Create)
		permissionRouter.GET("", permissionController.Search)
		permissionRouter.GET("all", permissionController.GetAll)
		permissionRouter.PATCH(":id", permissionController.Update)
		permissionRouter.DELETE(":id", permissionController.Delete)
	}
}
