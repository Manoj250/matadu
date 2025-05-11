package repo

import (
	"context"

	pb "matadu/cache/internal/pb"

	"go.uber.org/zap"
)

func (c *CacheRepo) RegisterServer(ctx context.Context, server *pb.ServerDetails) error {
	addr := server.ServerAddress
	if err := c.client.SAdd(ctx, "ws:servers", addr).Err(); err != nil {
		c.logger.Error("failed to add server", zap.Error(err))
		return err
	}
	return nil
}

func (c *CacheRepo) UnRegisterServer(ctx context.Context, server *pb.ServerDetails) error {
	addr := server.ServerAddress
	pipe := c.client.TxPipeline()
	pipe.SRem(ctx, "ws:servers", addr)
	pipe.Del(ctx, addr+"#active_users")
	if _, err := pipe.Exec(ctx); err != nil {
		c.logger.Error("failed to unregister server", zap.Error(err))
		return err
	}
	return nil
}

func (c *CacheRepo) GetServerLoad(ctx context.Context) ([]*pb.ServerLoad, error) {
	addrs, err := c.client.SMembers(ctx, "ws:servers").Result()
	if err != nil {
		c.logger.Error("failed to fetch server list", zap.Error(err))
		return nil, err
	}

	var loads []*pb.ServerLoad
	for _, addr := range addrs {
		val, err := c.client.Get(ctx, addr+"#active_users").Int()
		if err != nil && err.Error() != "redis: nil" {
			c.logger.Warn("failed to fetch user count for server", zap.String("addr", addr), zap.Error(err))
			continue
		}
		loads = append(loads, &pb.ServerLoad{
			ServerAddress: addr,
			Load:          int32(val),
		})
	}

	return loads, nil
}

func (c *CacheRepo) SetServerLoad(ctx context.Context, input []*pb.ServerLoad) error {
	pipe := c.client.TxPipeline()

	for _, s := range input {
		key := s.ServerAddress + "#active_users"
		pipe.Set(ctx, key, s.Load, 0)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		c.logger.Error("failed to set server loads", zap.Error(err))
		return err
	}
	return nil
}
