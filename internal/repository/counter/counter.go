package counter

import "github.com/redis/go-redis/v9"

const counterKey = "counter" // redis key for keeping counter

type Repo struct {
	redisClient *redis.Client
}

func NewRepo(redisClient *redis.Client) *Repo {
	return &Repo{redisClient: redisClient}
}
