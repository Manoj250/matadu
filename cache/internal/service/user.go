package service

import (
	"context"

	pb "matadu/cache/internal/pb"
)

func (c *CacheService) AddUserToCache(ctx context.Context, userDetails *pb.UserDetails) error {
	return c.repo.AddUserToCache(ctx, userDetails)
}

func (c *CacheService) RemoveUserFromCache(ctx context.Context, userId string) error {
	return c.repo.RemoveUserFromCache(ctx, userId)
}

func (c *CacheService) GetUserFromCache(ctx context.Context, userId string) (*pb.UserDetails, error) {
	return c.repo.GetUserFromCache(ctx, userId)
}
