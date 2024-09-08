package global

import (
	"database/sql"
	"github.com/open-auth/pkg/logger"
	"github.com/open-auth/pkg/setting"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	TokenPrivateKey    string = "PRIVATE_KEY"
	TokenPublicKey     string = "PUBLIC_KEY"
	AccessTokenExpire         = 3 * time.Hour
	RefreshTokenExpire        = 168 * time.Hour
	RefreshTokenKey    string = "x-rToken"
	BearerTokenKey     string = "Authorization"
)

var (
	Config setting.Config
	Logger *logger.Zap
	Mdb    *sql.DB
	Rdb    *redis.Client
)
