package user

import (
	"context"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Create(ctx context.Context, user models.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (last_name, first_name) VALUES ($1, $2) RETURNING id", userTable)

	var userID int
	if err := r.DB.QueryRowContext(ctx, query, user.LastName, user.FirstName).Scan(&userID); err != nil {
		return 0, fmt.Errorf("can not insert user into: %w", err)
	}

	return userID, nil
}
