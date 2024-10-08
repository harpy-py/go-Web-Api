package main

import (
	"github.com/harpy-py/go-Web-Api/api"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/data/cache"
	"github.com/harpy-py/go-Web-Api/data/db"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
)

// @
func main() {
	conf := config.GetConfig()
	logger := logging.NewLogger(conf)
	err := cache.InitRedis(conf)
	defer cache.CloseRedis()
	if err != nil{
		logger.Fatal(logging.Redis, logging.StartUp, err.Error(), nil)
	}

	err = db.InitDb(conf)
	defer db.CloseDb()
	if err != nil{
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}
	api.Initserver(conf)
}