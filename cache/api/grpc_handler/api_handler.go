package grpc_handler

import (
	pb "matadu/cache/internal/pb"
	"matadu/cache/internal/service"

	"go.uber.org/zap"
)

type CacheServiceServer struct {
	pb.UnimplementedCacheServiceServer
	service *service.CacheService
	logger  *zap.Logger
}

func NewCacheServiceServer(service *service.CacheService, zapLogger *zap.Logger) *CacheServiceServer {
	return &CacheServiceServer{
		service: service,
		logger:  zapLogger,
	}
}
