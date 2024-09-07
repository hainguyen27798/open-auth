//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/go-open-auth/internal/controllers"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/internal/services"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repos.NewUserRepo,
		services.NewUserService,
		controllers.NewUserController,
	)
	return new(controllers.UserController), nil
}
