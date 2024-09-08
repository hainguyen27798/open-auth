package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/middlewares"
	"github.com/go-open-auth/internal/wires"
)

type PermissionRouter struct{}

func (pr *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	permissionController, _ := wires.InitPermissionRouterHandler()

	permissionRouter := Router.Group("permissions")
	permissionRouter.Use(middlewares.AuthMiddleware())
	{
		permissionRouter.POST("", permissionController.Create)
		permissionRouter.GET("", permissionController.GetAll)
		permissionRouter.PATCH(":id", permissionController.Update)
		permissionRouter.DELETE(":id", permissionController.Delete)
	}
}
