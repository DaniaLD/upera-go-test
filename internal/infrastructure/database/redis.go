package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(uri string) (client *redis.Client, err error) {
	opts, err := redis.ParseURL(uri)
	client = redis.NewClient(opts)
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("\"Error pinging Redis: %v", err)
	}
	return client, nil
}
