package service

import (
	"context"
	
	pb "matadu/cache/internal/pb"
)

func (c *CacheService) AddUserToChat(ctx context.Context, userId, chatId string) error {
	return c.repo.AddUserToChat(ctx, userId, chatId)
}

func (c *CacheService) RemoveUserFromChat(ctx context.Context, userId, chatId string) error {
	return c.repo.RemoveUserFromChat(ctx, userId, chatId)
}

func (c *CacheService) GetUsersByChatId(ctx context.Context, chatId string) ([]*pb.UserDetails, error) {
	return c.repo.GetUsersByChatId(ctx, chatId)
}
