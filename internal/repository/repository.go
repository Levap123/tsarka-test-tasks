package repository

import (
	"context"
	"database/sql"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
	"github.com/Levap123/tsarka-test-tasks/internal/repository/counter"
	"github.com/Levap123/tsarka-test-tasks/internal/repository/user"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	CounterRepoI
	UserRepoI
}

type CounterRepoI interface {
	Get(ctx context.Context) (int, error)
	Set(ctx context.Context, res int) error
}

type UserRepoI interface {
	Create(ctx context.Context, user models.User) (int, error)
	Get(ctx context.Context, userID int) (models.User, error)
	Delete(ctx context.Context, userID int) error
	Update(ctx context.Context, user models.User) error
}

func NewRepository(cl *redis.Client, DB *sql.DB) *Repository {
	return &Repository{
		CounterRepoI: counter.NewRepo(cl),
		UserRepoI:    user.NewRepo(DB),
	}
}
