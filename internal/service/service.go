package service

import (
	"context"
	"sync"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/matchmaker/internal/repository"

	"github.com/BelyaevEI/platform_common/pkg/db"
)

// UserServicer entity for service layer
type UserServicer interface {
	AddUserToPool(ctx context.Context, body []byte) error
	FindPalyers(ctx context.Context) ([]model.User, error)
	PrintNewGroup(users []model.User)
}

// Struct implementation api layer
type service struct {
	txManager      db.TxManager
	userRepository repository.UserRepositorer
	countGroup     int32
	mutex          sync.Mutex
}

// NewService Constructor
func NewService(userRepository repository.UserRepositorer,
	txManager db.TxManager) UserServicer {
	return &service{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
