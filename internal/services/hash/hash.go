package hash

import "github.com/Levap123/tsarka-test-tasks/internal/repository"

type HashService struct {
	repo repository.HashRepoI
}

func NewService(repo repository.HashRepoI) *HashService {
	return &HashService{
		repo: repo,
	}
}
