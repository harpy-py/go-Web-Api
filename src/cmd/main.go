package main

import (
	"log"

	"github.com/harpy-py/go-Web-Api/api"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/data/cache"
	"github.com/harpy-py/go-Web-Api/data/db"
)

func main() {
	conf := config.GetConfig()
	err := cache.InitRedis(conf)
	defer cache.CloseRedis()
	if err != nil{
		log.Fatal(err)
	}

	err = db.InitDb(conf)
	defer db.CloseDb()
	if err != nil{
		log.Fatal(err)
	}
	api.Initserver(conf)
}