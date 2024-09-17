package initialize

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"time"
)

func CheckErrorPanic(err error, errMsg string) {
	if err != nil {
		global.Logger.Error(errMsg, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	config := global.Config.Mysql

	// sql config connection
	cfg := mysql.Config{
		User:      config.Username,
		Passwd:    config.Password,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", config.Host, config.Port),
		DBName:    config.Database,
		ParseTime: true,
		Loc:       time.Local,
	}
	dsn := cfg.FormatDSN()
	db, _ := sql.Open("mysql", dsn)
	err := db.Ping()
	CheckErrorPanic(err, "Init mysql failed")
	global.Logger.Info("Initializing mysql successfully")

	setPool(db)
	migrateTables(db)

	global.Mdb = sqlx.NewDb(db, "mysql")
}

func setPool(db *sql.DB) {
	config := global.Config.Mysql

	db.SetConnMaxIdleTime(time.Duration(config.MaxIdleConn))
	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime))
}

func migrateTables(db *sql.DB) {
	if err := goose.SetDialect("mysql"); err != nil {
		CheckErrorPanic(err, "Init mysql failed")
	}

	if err := goose.Up(db, "database/migrations"); err != nil {
		CheckErrorPanic(err, "Init mysql failed")
	}
}
