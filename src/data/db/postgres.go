package db

import (
	"fmt"
	"time"

	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDb(conf *config.Config) error {
	var err error
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.DbName,
		conf.Postgres.SslMode)

	dbClient, err = gorm.Open(postgres.Open(connect), &gorm.Config{})
	if err != nil{
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil{
		return err
	}

	sqlDb.SetMaxIdleConns(conf.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(conf.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(conf.Postgres.ConnMaxLifetime *time.Minute)

	logger.Info(logging.Postgres, logging.StartUp, "The Connection has established!", nil)
	return nil
}

func GetDb() *gorm.DB{
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}