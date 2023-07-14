package counter

import "context"

func (s *CounterService) Add(ctx context.Context, add int) error {
	val, err := s.repo.Get(ctx)
	if err != nil {
		return err
	}
	val += add
	return s.repo.Set(ctx, val)
}
