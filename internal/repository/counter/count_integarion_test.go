package counter

import (
	"context"
	"log"
	"testing"

	"github.com/Levap123/tsarka-test-tasks/internal/infrastructure/redis"
)

var counterRepo *Repo

func TestMain(m *testing.M) {
	redisCl, err := redis.New()
	if err != nil {
		log.Fatalf("can not create redis client: %v", err)
	}

	counterRepo = NewRepo(redisCl)
	m.Run()
}

func Test_Count(t *testing.T) {
	t.Run("should work fine", func(t *testing.T) {
		val, err := counterRepo.Get(context.Background())
		if err != nil {
			t.Errorf("expected noerr, got: %v", err)
		}

		if err := counterRepo.Set(context.Background(), val+10); err != nil {
			t.Errorf("expected noerr, got: %v", err)
		}

		valActual, err := counterRepo.Get(context.Background())
		if err != nil {
			t.Errorf("expected noerr, got: %v", err)
		}

		if valActual != val+10 {
			t.Errorf("expected equal, got: %d %d", val, valActual)
		}
	})
	t.Run("should return zero", func(t *testing.T) {
		if err := counterRepo.Set(context.Background(), 0); err != nil {
			t.Errorf("expected noerr, got: %v", err)
		}

		val, err := counterRepo.Get(context.Background())
		if err != nil {
			t.Errorf("expected noerr, got: %v", err)
		}

		if val != 0 {
			t.Errorf("expected 0z, got: %d", val)
		}
	})
}
