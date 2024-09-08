//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/go-open-auth/internal/controllers"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/internal/services"
	"github.com/google/wire"
)

func InitPermissionRouterHandler() (*controllers.PermissionController, error) {
	wire.Build(
		repos.NewPermissionRepo,
		services.NewPermissionService,
		controllers.NewPermissionController,
	)
	return new(controllers.PermissionController), nil
}
