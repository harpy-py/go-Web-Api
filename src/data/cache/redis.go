package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/harpy-py/go-Web-Api/config"
)

var redisClient *redis.Client

func InitRedis(conf *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Password,
		DB: 0,
		DialTimeout: conf.Redis.DialTimeout * time.Second,
		ReadTimeout: conf.Redis.ReadTimeout * time.Second,
		WriteTimeout: conf.Redis.WriteTimeout * time.Second,
		PoolSize: conf.Redis.PoolSize,
		PoolTimeout: conf.Redis.PoolTimeout,
		IdleTimeout: conf.Redis.IdleTimeout * time.Millisecond,
		IdleCheckFrequency: conf.Redis.IdleCheckFrequency * time.Millisecond,
	})

	_, err := redisClient.Ping().Result()
	if err != nil{
		return err
	}
	return nil
}

func GetRedis() *redis.Client{
	return redisClient
}

func CloseRedis(){
	redisClient.Close()
}