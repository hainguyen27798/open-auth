package initialize

import (
	"database/sql"
	"fmt"
	"github.com/go-open-auth/global"
	_ "github.com/go-sql-driver/mysql"
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
	dsnFormat := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		dsnFormat,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, _ := sql.Open("mysql", dsn)
	err := db.Ping()
	CheckErrorPanic(err, "Init mysql failed")
	global.Logger.Info("Initializing mysql successfully")
	global.Mdb = db

	setPool()
	migrateTables()
}

func setPool() {
	config := global.Config.Mysql
	sqlDB := global.Mdb

	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleConn))
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime))
}

func migrateTables() {
	db := global.Mdb

	if err := goose.SetDialect("mysql"); err != nil {
		CheckErrorPanic(err, "Init mysql failed")
	}

	if err := goose.Up(db, "database/migrations"); err != nil {
		CheckErrorPanic(err, "Init mysql failed")
	}
}
