package repository

import (
	"context"

	"github.com/BelyaevEI/platform_common/pkg/db"
)

type UserRepositorer interface {
	SearchMatch(ctx context.Context) error
}

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) UserRepositorer {
	return &repository{db: db}
}
