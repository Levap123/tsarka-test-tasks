package counter

import "context"

func (s *CounterService) Sub(ctx context.Context, sub int) error {
	val, err := s.repo.Get(ctx)
	if err != nil {
		return err
	}
	val -= sub
	return s.repo.Set(ctx, val)
}
