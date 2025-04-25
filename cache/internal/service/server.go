package service

import (
	"context"

	pb "matadu/cache/internal/pb"
)

func (c *CacheService) RegisterServer(ctx context.Context, server *pb.ServerDetails) error {
	return c.repo.RegisterServer(ctx, server)
}

func (c *CacheService) UnRegisterServer(ctx context.Context, server *pb.ServerDetails) error {
	return c.repo.UnRegisterServer(ctx, server)
}

func (c *CacheService) GetServerLoad(ctx context.Context) ([]*pb.ServerLoad, error) {
	return c.repo.GetServerLoad(ctx)
}

func (c *CacheService) SetServerLoad(ctx context.Context, input []*pb.ServerLoad) error {
	return c.repo.SetServerLoad(ctx, input)
}
