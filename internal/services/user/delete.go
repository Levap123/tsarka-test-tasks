package user

import "context"

func (s *UserService) Delete(ctx context.Context, userID int) error {
	return s.repo.Delete(ctx, userID)
}
