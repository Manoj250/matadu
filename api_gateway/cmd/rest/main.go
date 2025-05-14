package main

import (
	"api_gateway/api/rest"
	db "api_gateway/internal/db/redis"
	"api_gateway/internal/repo/redis"
	"api_gateway/internal/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load()

	var logger *zap.Logger
	var err error

	if os.Getenv("APP_ENV") == "PRODUCTION" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	redisClient := db.NewRedisClusterClient()
	repo := redis.NewRepo(redisClient, logger)
	svc := service.NewService(repo, logger)

	router := gin.Default()

	router.Use(rest.ZapLoggerMiddleware(logger), gin.Recovery())

	//Inject service & logger into handler
	rest.RegisterRoutes(router, svc, logger)

	logger.Info("REST API running on :8080")
	_ = router.Run(":8080")
}
