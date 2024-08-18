package repository

import (
	"context"
	"log"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/platform_common/pkg/db"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) AddUserToPool(ctx context.Context, user model.User) error {

	if r.storageFlag {

		// if pool not init, need initializating
		if r.poolUsers == nil {
			r.mutex.Lock()
			r.poolUsers = make([]model.User, 0)
			r.mutex.Unlock()
		}

		// Add user to memory storage
		r.mutex.Lock()
		defer r.mutex.Unlock()

		r.poolUsers = append(r.poolUsers, user)
	} else {
		if err := r.addUserToPoolDB(ctx, user); err != nil {
			log.Printf("insert user to db is failed")
			return err
		}
	}

	return nil
}

func (r *repository) addUserToPoolDB(ctx context.Context, user model.User) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, skillColumn, latencyColumn).
		Values(user.Name, user.Skill, user.Latency)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "add user",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
