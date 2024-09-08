package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/router"
	"os"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if os.Getenv("MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// Declare router
	authRouter := router.AppRouter.Auth
	adminRouter := router.AppRouter.Admin

	mainRouter := r.Group("/v1")
	{
		authRouter.InitAuthenticateRouter(mainRouter)
	}
	{
		adminRouter.InitUserRouter(mainRouter)
		adminRouter.InitPermissionRouter(mainRouter)
		adminRouter.InitRoleRouter(mainRouter)
	}

	return r
}
