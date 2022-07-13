package config

import (
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func BuildRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
