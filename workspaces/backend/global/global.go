package global

import (
	"database/sql"
	"github.com/go-open-auth/pkg/logger"
	"github.com/go-open-auth/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.Zap
	Mdb    *sql.DB
	Rdb    *redis.Client
)
