package initialize

import "github.com/open-auth/internal/services"

func SetupAdmin() {
	configService := services.NewConfigService()
	configService.InitAdmin("admin@auth.com", "auth@123")
}
