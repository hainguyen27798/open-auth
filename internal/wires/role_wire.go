//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/open-auth/internal/controllers"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/internal/services"
)

func InitRoleRouterHandler() (*controllers.RoleController, error) {
	wire.Build(
		repos.NewRoleRepo,
		services.NewRoleService,
		controllers.NewRoleController,
	)
	return new(controllers.RoleController), nil
}
