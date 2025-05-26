package redis

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClusterClient() *redis.ClusterClient {
	addrs := strings.Split(os.Getenv("REDIS_CLUSTER_ADDRS"), ",")

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        addrs,
		// Password:     os.Getenv("REDIS_PASS"), // optional
		MaxRetries:   3,
		PoolSize:     50, // number of connections per node
		MinIdleConns: 10, // keep-alive connections
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// Health check
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("REDIS IS DEAD: " + err.Error())
	}

	return rdb
}
