package user

import (
	"context"
	"fmt"
)

func (r *Repo) Delete(ctx context.Context, ID int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1", userTable)

	if _, err := r.DB.ExecContext(ctx, query, ID); err != nil {
		return err
	}

	return nil
}
