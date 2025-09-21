package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisInit(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("failed to connect to redis '%s': %v", Config.Redis.Addr, err)
		return nil, err
	}

	log.Printf("Successfully connected to redis: %s", pong)
	return client, nil
}
