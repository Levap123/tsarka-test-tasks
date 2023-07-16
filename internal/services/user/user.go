package user

import "github.com/Levap123/tsarka-test-tasks/internal/repository"

type UserService struct {
	repo repository.UserRepoI
}

func NewService(repo repository.UserRepoI) *UserService {
	return &UserService{
		repo: repo,
	}
}
