package repository

import (
	"context"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) DeleteUsers(ctx context.Context, users []model.User) error {
	if r.storageFlag {
		return r.deleteUsersMem()
	}

	return r.deleteUsersDB(ctx, users)
}

func (r *repository) deleteUsersMem() error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	// slice has been sorted
	r.poolUsers = r.poolUsers[r.groupsize+1:]

	return nil
}

func (r *repository) deleteUsersDB(ctx context.Context, users []model.User) error {

	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar)

	for _, v := range users {
		builder = builder.Where(
			sq.And{
				sq.Eq{"name": v.Name},
				sq.Eq{"skill": v.Skill},
				sq.Eq{"latency": v.Latency},
			},
		)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "delete user",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
