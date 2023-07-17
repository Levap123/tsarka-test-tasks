package hash

import (
	"context"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Create(ctx context.Context, hash models.Hash) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (status, hash) VALUES ($1, $2) RETURNING id", hashTable)

	var hashID int
	if err := r.DB.QueryRowContext(ctx, query, hash.Status, hash.Hash).Scan(&hashID); err != nil {
		return 0, fmt.Errorf("can not insert into hash_records: %w", err)
	}

	return hashID, nil
}
