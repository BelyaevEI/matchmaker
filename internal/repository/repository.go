package repository

import (
	"context"
	"sync"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/platform_common/pkg/db"
)

const (
	tableName     = "user"
	nameColumn    = "name"
	skillColumn   = "skill"
	latencyColumn = "latency"
)

type UserRepositorer interface {
	AddUserToPool(ctx context.Context, user model.User) error
}

// Struct implementation service layer
type repository struct {
	db          db.Client
	poolUsers   []model.User
	storageFlag bool // true - memory; false - db;
	mutex       sync.RWMutex
}

func NewRepository(db db.Client, storageFlag bool) UserRepositorer {

	return &repository{
		db:          db,
		storageFlag: storageFlag,
	}
}
