package repository

import (
	"context"

	"github.com/Levap123/tsarka-test-tasks/internal/repository/counter"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	CounterRepoI
}

type CounterRepoI interface {
	Get(ctx context.Context) (int, error)
	Set(ctx context.Context, res int) error
}

func NewRepository(cl *redis.Client) *Repository {
	return &Repository{
		CounterRepoI: counter.NewRepo(cl),
	}
}
