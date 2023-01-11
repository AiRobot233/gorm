package config

import (
	"gin/utils"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func BuildRedis() *redis.Client {
	redisDb := utils.GetEnvData("REDIS_DB")
	if redisDb == "" {
		redisDb = "0"
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.GetEnvData("REDIS_ADDR"),
		Password: utils.GetEnvData("REDIS_PW"),
		DB:       utils.StrToInt(redisDb),
	})
	return rdb
}
