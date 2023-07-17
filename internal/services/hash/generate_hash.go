package hash

import (
	"context"
	"hash/crc64"
	"log"
	"time"

	"github.com/Levap123/tsarka-test-tasks/internal/models"
)

var table = crc64.MakeTable(crc64.ISO)

func (s *HashService) GenerateHash(ctx context.Context, val string) (int, error) {
	hashSum := crc64.Checksum([]byte(val), table)

	hash := models.Hash{
		Status: models.PendindStatus,
	}

	id, err := s.repo.Create(ctx, hash)
	if err != nil {
		log.Printf("can not create hash: %v\n", err)
		return 0, err
	}

	ctx = context.Background()
	go func() {
		hash.ID = id

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		finish := time.After(60 * time.Second)

		for {
			select {
			case <-ticker.C:
				hashSum = hashSum & uint64(time.Now().UnixNano())
			case <-finish:
				hash.Hash = s.getOnes(hashSum)
				hash.Status = models.FinishedStatus

				if err := s.repo.Update(ctx, hash); err != nil {
					log.Printf("can not update hash: %v\n", err)
				}
				return
			}
		}
	}()

	return id, nil
}

func (s *HashService) getOnes(hashSum uint64) int {
	count := 0

	for hashSum > 0 {
		count += int(hashSum & 1)
		hashSum >>= 1
	}

	return count
}
