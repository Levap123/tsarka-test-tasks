package counter

import "context"

// no need for tests
func (s *CounterService) Get(ctx context.Context) (int, error) {
	return s.repo.Get(ctx)
}
