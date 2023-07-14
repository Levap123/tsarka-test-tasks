package counter

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *Repo) Get(ctx context.Context) (int, error) {
	val, err := r.redisClient.Get(ctx, counterKey).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return valInt, nil
}

// expiration hardcoded because im too lazy for configs
func (r *Repo) Set(ctx context.Context, res int) error {
	return r.redisClient.Set(ctx, counterKey, res, time.Minute*15).Err()
}
