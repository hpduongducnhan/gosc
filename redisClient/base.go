package redisclient

import (
	"context"

	"github.com/rs/zerolog/log"

	redis "github.com/redis/go-redis/v9"
)

func ConnectRedis(ctx context.Context, url string) (*redis.Client, error) {
	// sample
	// url = "redis://user:password@localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opts)

	res, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to Redis: %s", res)

	return client, nil
}
