package hash

import (
	"context"
	"log"
	"strconv"
)

func (s *HashService) Get(ctx context.Context, ID int) (string, error) {
	hash, err := s.repo.Get(ctx, ID)
	if err != nil {
		log.Printf("can not get hash: %v\n", err)
		return "", err
	}

	if hash.Status == "" {
		return "not found", nil
	}

	if hash.IsPending() {
		return "в обработке", nil
	}

	return strconv.Itoa(hash.Hash), nil
}
