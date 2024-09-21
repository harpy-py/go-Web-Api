package main

import (
	"github.com/harpy-py/go-Web-Api/api"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/data/cache"
)

func main() {
	conf := config.GetConfig()
	cache.InitRedis(conf)
	defer cache.CloseRedis()
	api.Initserver(conf)
}