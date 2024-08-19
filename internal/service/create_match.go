package service

import (
	"context"

	"github.com/BelyaevEI/matchmaker/internal/model"
)

// Find players for create match
func (s *service) FindPalyers(ctx context.Context) ([]model.User, error) {

	var users []model.User

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {

		// Find user with older time in pool
		userOld, errTx := s.userRepository.FindOldUser(ctx)
		if errTx != nil {
			return errTx
		}

		// Find other user for old user
		users, errTx = s.userRepository.FindUsersForMatch(ctx, userOld)
		if errTx != nil {
			return errTx
		}

		// Delete finder users
		errTx = s.userRepository.DeleteUsers(ctx, users)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}
