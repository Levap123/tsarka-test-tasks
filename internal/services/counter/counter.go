package counter

import "github.com/Levap123/tsarka-test-tasks/internal/repository"

type CounterService struct {
	repo repository.CounterRepoI
}

func NewService(repo repository.CounterRepoI) *CounterService {
	return &CounterService{repo: repo}
}
