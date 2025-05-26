package service

import (
	"api_gateway/internal/repo/redis"

	"go.uber.org/zap"
)

type Service struct {
	repo   *redis.Repo
	logger *zap.Logger
}

func NewService(repo *redis.Repo, zapLogger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: zapLogger,
	}
}
