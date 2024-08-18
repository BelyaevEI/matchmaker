package service

import (
	"context"

	"github.com/BelyaevEI/matchmaker/internal/repository"
)

type UserServicer interface {
	SearchMatch(ctx context.Context, body []byte) error
}

// Struct implementation api layer
type service struct {
	userRepository repository.UserRepositorer
	groupSize      int32
}

// Constructor
func NewService(userRepository repository.UserRepositorer, groupSize int32) UserServicer {
	return &service{
		userRepository: userRepository,
		groupSize:      groupSize,
	}
}
