package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// hardcoded because i too lazy for adding configs
func New() (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := c.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return c, nil
}
