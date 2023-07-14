package services

import (
	"context"

	"github.com/Levap123/tsarka-test-tasks/internal/repository"
	"github.com/Levap123/tsarka-test-tasks/internal/services/algos"
	"github.com/Levap123/tsarka-test-tasks/internal/services/counter"
)

type Services struct {
	AlgosServiceI
	CounterServiceI
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

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		AlgosServiceI:   algos.NewService(),
		CounterServiceI: counter.NewService(repo.CounterRepoI),
	}
}
