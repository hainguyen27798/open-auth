package initialize

import (
	"github.com/go-open-auth/global"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {
	name := os.Getenv("MODE")

	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName(name)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
