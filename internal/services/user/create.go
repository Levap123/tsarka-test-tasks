package user

import (
	"context"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (s *UserService) Create(ctx context.Context, firstName, lastName string) (int, error) {
	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
	}
	return s.repo.Create(ctx, user)
}
