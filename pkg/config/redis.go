package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisInit(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8081",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to redis: %s", pong)
	return client, nil
}
