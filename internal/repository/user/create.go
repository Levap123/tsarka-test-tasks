package user

import (
	"context"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Create(ctx context.Context, user models.User) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (last_name, first_name) VALUES ($1, $2)", userTable)

	res, err := r.DB.ExecContext(ctx, query, user.LastName, user.FirstName)
	if err != nil {
		return 0, fmt.Errorf("can not insert user into: %w", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("can not get user id: %w", err)
	}

	return ID, nil
}
