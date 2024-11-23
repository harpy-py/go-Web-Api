package migrations

import (
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/data/db"
	"github.com/harpy-py/go-Web-Api/data/models"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country){
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city){
		tables = append(tables, city)
	}
	
	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Tables created!", nil)
}

func Down_1() {

}