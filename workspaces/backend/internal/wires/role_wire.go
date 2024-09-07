//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/go-open-auth/internal/controllers"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/internal/services"
	"github.com/google/wire"
)

func InitRoleRouterHandler() (*controllers.RoleController, error) {
	wire.Build(
		repos.NewRoleRepo,
		services.NewRoleService,
		controllers.NewRoleController,
	)
	return new(controllers.RoleController), nil
}
