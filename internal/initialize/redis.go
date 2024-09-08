package initialize

import (
	"context"
	"fmt"
	"github.com/go-open-auth/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	config := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("redis init error:", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Initializing redis successfully")
	global.Rdb = rdb
}
