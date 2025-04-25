package repo

import (
	"context"
	pb "matadu/cache/internal/pb"

	redis "github.com/redis/go-redis/v9"
)

func (c *CacheRepo) AddUserToCache(ctx context.Context, userDetails *pb.UserDetails) error {
	return c.client.HSet(ctx, "user:"+userDetails.UserId,
		"id", userDetails.UserId,
		"name", userDetails.UserName,
		"email", userDetails.UserEmail,
		"phone", userDetails.UserPhone,
		"status", userDetails.UserStatus,
		"device_id", userDetails.DeviceId,
	).Err()
}

func (c *CacheRepo) RemoveUserFromCache(ctx context.Context, userId string) error {
	return c.client.Del(ctx, "user:"+userId).Err()
}

func (c *CacheRepo) GetUserFromCache(ctx context.Context, userId string) (*pb.UserDetails, error) {
	data, err := c.client.HGetAll(ctx, "user:"+userId).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // user not found
		}
		return nil, err
	}

	// Construct the user details from the hash fields
	user := &pb.UserDetails{
		UserId:     data["id"],
		UserName:   data["name"],
		UserEmail:  data["email"],
		UserPhone:  data["phone"],
		UserStatus: data["status"],
		DeviceId:   data["device_id"],
	}

	return user, nil
}
