package repository

import (
	"context"
	"sync"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/platform_common/pkg/db"
)

const (
	tableName       = "user"
	nameColumn      = "name"
	skillColumn     = "skill"
	latencyColumn   = "latency"
	createdAtColumn = "created_at"
)

type UserRepositorer interface {
	AddUserToPool(ctx context.Context, user model.User) error
	FindOldUser(ctx context.Context) (model.User, error)
	FindUsersForMatch(ctx context.Context, userOld model.User) ([]model.User, error)
	DeleteUsers(ctx context.Context, users []model.User) error
}

// Struct implementation service layer
type repository struct {
	db          db.Client
	poolUsers   []model.User
	storageFlag bool // true - memory; false - db;
	groupsize   int32
	mutex       sync.RWMutex
}

func NewRepository(db db.Client, storageFlag bool, groupSize int32) UserRepositorer {

	return &repository{
		db:          db,
		storageFlag: storageFlag,
		groupsize:   groupSize,
	}
}
