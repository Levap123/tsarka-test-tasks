package user

import "context"

func (s *UserService) Update(ctx context.Context, userID int, firstName, lastName string) error {
	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return err
	}

	if firstName != "" {
		user.FirstName = firstName
	}

	if lastName != "" {
		user.LastName = lastName
	}

	return s.repo.Update(ctx, user)
}
