package redis

import (
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Repo struct {
	logger *zap.Logger
	client *redis.ClusterClient
}

func NewRepo(redisClusterClient *redis.ClusterClient, zapLogger *zap.Logger) *Repo {
	return &Repo{
		client: redisClusterClient,
		logger: zapLogger,
	}
}