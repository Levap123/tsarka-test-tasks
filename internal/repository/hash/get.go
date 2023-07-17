package hash

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Get(ctx context.Context, ID int) (models.Hash, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", hashTable)

	var hash models.Hash
	if err := r.DB.QueryRowContext(ctx, query, ID).Scan(&hash.ID, &hash.Status, &hash.Hash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Hash{}, nil
		}
		return models.Hash{}, fmt.Errorf("can not get hash_record: %w", err)
	}

	return hash, nil
}
