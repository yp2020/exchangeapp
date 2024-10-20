package config

import (
	"exchangeapp/global"
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GlobalConfig.RedisConfig.Addr,
		Password: GlobalConfig.RedisConfig.Password, // no password set
		DB:       GlobalConfig.RedisConfig.DB,       // use default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to configure redis, got error: %v", err)
	}
	global.RedisDB = RedisClient
}
