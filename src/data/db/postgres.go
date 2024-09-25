package db

import (
	"fmt"
	"log"
	"time"

	"github.com/harpy-py/go-Web-Api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(conf *config.Config) error {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.DbName,
		conf.Postgres.SslMode)

	dbClient, err := gorm.Open(postgres.Open(connect), &gorm.Config{})
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

	log.Println("The Connection has established!")
	return nil
}

func GetDb() *gorm.DB{
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}