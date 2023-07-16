package services

import (
	"context"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
	"github.com/Levap123/tsarka-test-tasks/internal/repository"
	"github.com/Levap123/tsarka-test-tasks/internal/services/algos"
	"github.com/Levap123/tsarka-test-tasks/internal/services/counter"
	"github.com/Levap123/tsarka-test-tasks/internal/services/user"
)

type Services struct {
	AlgosServiceI
	CounterServiceI
	UserServiceI
}

type AlgosServiceI interface {
	FindSubstring(str string) string
	EmailCheck(str string) []string
}

type CounterServiceI interface {
	Add(ctx context.Context, add int) error
	Sub(ctx context.Context, sub int) error
	Get(ctx context.Context) (int, error)
}

type UserServiceI interface {
	Create(ctx context.Context, firstName, lastName string) (int, error)
	Get(ctx context.Context, userID int) (models.User, error)
	Delete(ctx context.Context, userID int) error
	Update(ctx context.Context, userID int, firstName, lastName string) error
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		AlgosServiceI:   algos.NewService(),
		CounterServiceI: counter.NewService(repo.CounterRepoI),
		UserServiceI:    user.NewService(repo.UserRepoI),
	}
}
