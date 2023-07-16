package user

import (
	"context"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Update(ctx context.Context, user models.User) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1, lastname = $2 WHERE id = $3", userTable)

	if _, err := r.DB.ExecContext(ctx, query, user.FirstName, user.LastName, user.ID); err != nil {
		return fmt.Errorf("can not update user: %w", err)
	}

	return nil
}
