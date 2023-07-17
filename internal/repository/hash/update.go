package hash

import (
	"context"
	"fmt"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

func (r *Repo) Update(ctx context.Context, hash models.Hash) error {
	query := fmt.Sprintf("UPDATE %s SET status=$1, hash=$2 WHERE id=$3", hashTable)

	if _, err := r.DB.ExecContext(ctx, query, hash.Status, hash.Hash, hash.ID); err != nil {
		return fmt.Errorf("can not update hash_record: %w", err)
	}

	return nil
}
