package repo

import (
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type CacheRepo struct {
	logger *zap.Logger
	client *redis.Client
}

func NewCacheRepo(redisClient *redis.Client, zapLogger *zap.Logger) *CacheRepo {
	return &CacheRepo{
		client: redisClient,
		logger: zapLogger,
	}
}
