package initialize

import (
	"fmt"
	"github.com/open-auth/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	err := r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
	if err != nil {
		return
	}
}
