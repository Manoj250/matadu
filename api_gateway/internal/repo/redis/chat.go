package repo

import (
	"context"

	pb "matadu/cache/internal/pb"
)

func (c *CacheRepo) AddUserToChat(ctx context.Context, userId, chatId string) error {
	key := "chat:" + chatId + ":users"
	return c.client.SAdd(ctx, key, userId).Err()
}

func (c *CacheRepo) RemoveUserFromChat(ctx context.Context, userId, chatId string) error {
	key := "chat:" + chatId + ":users"
	return c.client.SRem(ctx, key, userId).Err()
}

func (c *CacheRepo) GetUsersByChatId(ctx context.Context, chatId string) ([]*pb.UserDetails, error) {
	key := "chat:" + chatId + ":users"
	userIds, err := c.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var users []*pb.UserDetails
	for _, userId := range userIds {
		userDetails, err := c.GetUserFromCache(ctx, userId)
		if err != nil {
			return nil, err
		}

		if userDetails == nil {
			continue
		}

		users = append(users, userDetails)
	}

	return users, nil
}

