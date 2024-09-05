//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/go-open-auth/internal/controllers"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/internal/services"
	"github.com/google/wire"
)

func InitAuthRouterHandler() (*controllers.AuthController, error) {
	wire.Build(
		repos.NewUserRepo,
		repos.NewUserAuthRepo,
		services.NewAuthService,
		controllers.NewAuthController,
	)
	return new(controllers.AuthController), nil
}
