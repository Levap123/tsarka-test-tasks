package user

import (
	"context"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (s *UserService) Get(ctx context.Context, userID int) (models.User, error) {
	return s.repo.Get(ctx, userID)
}
