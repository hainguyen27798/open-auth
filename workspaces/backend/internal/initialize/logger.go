package initialize

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
