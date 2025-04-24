package main

import (
	"errors"
	"log"
	"net"
	"os"

	grpcHandler "matadu/cache/api/grpc_handler"
	"matadu/cache/internal/db"
	pb "matadu/cache/internal/pb"
	"matadu/cache/internal/repo"
	"matadu/cache/internal/service"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	var logger *zap.Logger

	err := godotenv.Load()

	if os.Getenv("APP_ENV") == "PRODUCTION" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	defer logger.Sync()

	err = run(logger)
	if err != nil {
		logger.Error("error during run", zap.Error(err))
	}
}

func run(zapLogger *zap.Logger) error {

	redisClient := db.NewRedisClient()
	if redisClient == nil {
		return errors.New("failed to connect to Redis")
	}

	repo := repo.NewCacheRepo(redisClient, zapLogger)
	service := service.NewCacheService(repo, zapLogger)
	handler := grpcHandler.NewCacheServiceServer(service, zapLogger)
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	pb.RegisterCacheServiceServer(grpcServer, handler)
	zapLogger.Info("gRPC server on :50051")
	return grpcServer.Serve(lis)
}
