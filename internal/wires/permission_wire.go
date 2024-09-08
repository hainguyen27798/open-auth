//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/open-auth/internal/controllers"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/internal/services"
)

func InitPermissionRouterHandler() (*controllers.PermissionController, error) {
	wire.Build(
		repos.NewPermissionRepo,
		services.NewPermissionService,
		controllers.NewPermissionController,
	)
	return new(controllers.PermissionController), nil
}
