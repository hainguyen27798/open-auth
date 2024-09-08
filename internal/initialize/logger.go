package initialize

import (
	"github.com/open-auth/global"
	"github.com/open-auth/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
