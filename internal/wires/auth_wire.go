//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/open-auth/internal/controllers"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/internal/services"
)

func InitAuthRouterHandler() (*controllers.AuthController, error) {
	wire.Build(
		repos.NewUserRepo,
		repos.NewUserAuthRepo,
		repos.NewTokenRepo,
		services.NewTokenService,
		services.NewAuthService,
		controllers.NewAuthController,
	)
	return new(controllers.AuthController), nil
}
