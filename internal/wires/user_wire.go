//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/open-auth/internal/controllers"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repos.NewUserRepo,
		services.NewUserService,
		controllers.NewUserController,
	)
	return new(controllers.UserController), nil
}
