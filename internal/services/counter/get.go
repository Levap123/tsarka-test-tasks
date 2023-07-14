package counter

import "context"

func (s *CounterService) Get(ctx context.Context) (int, error) {
	return s.repo.Get(ctx)
}
