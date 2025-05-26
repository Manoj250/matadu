package service

import (
	"matadu/cache/internal/repo"

	"go.uber.org/zap"
)

type CacheService struct {
	repo   *repo.CacheRepo
	logger *zap.Logger
}

func NewCacheService(repo *repo.CacheRepo, zapLogger *zap.Logger) *CacheService {
	return &CacheService{
		repo:   repo,
		logger: zapLogger,
	}
}
