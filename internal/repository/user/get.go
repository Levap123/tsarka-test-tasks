package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Get(ctx context.Context, ID int) (models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)

	var user models.User
	if err := r.DB.QueryRowContext(ctx, query, ID).
		Scan(user.ID, user.FirstName, user.LastName); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, nil
		}

		return models.User{}, err
	}

	return user, nil
}
