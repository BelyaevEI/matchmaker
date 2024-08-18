package service

import (
	"context"

	"github.com/BelyaevEI/matchmaker/internal/repository"
)

type UserServicer interface {
	SearchMatch(ctx context.Context) error
}

// Struct implementation api layer
type service struct {
	userRepository repository.UserRepositorer
}

// Constructor
func NewService(userRepository repository.UserRepositorer) UserServicer {
	return &service{userRepository: userRepository}
}
