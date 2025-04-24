package db

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"), 
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
}
