package repository

import (
	"context"
	"time"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/platform_common/pkg/db"

	sq "github.com/Masterminds/squirrel"
)

// Find user with older time in queue
func (r *repository) FindOldUser(ctx context.Context) (model.User, error) {

	// Switch mem or db
	// I know that maybe better create two repo
	if r.storageFlag {
		return r.findOldUserMem(ctx)
	} else {
		return r.findOldUserDB(ctx)
	}

}

func (r *repository) findOldUserMem(_ context.Context) (model.User, error) {

	var (
		user   model.User
		oldest = time.Now()
	)

	for _, v := range r.poolUsers {
		if len(user.Name) == 0 {
			user = v
		}

		if user.TimeQueue.Before(oldest) {
			oldest = user.TimeQueue
			user = v
		}
	}

	return user, nil
}

func (r *repository) findOldUserDB(ctx context.Context) (model.User, error) {
	var user model.User

	builder := sq.Select(nameColumn, skillColumn, latencyColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		OrderBy("created_at ASC").
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return model.User{}, err
	}

	q := db.Query{
		Name:     "select old user",
		QueryRaw: query,
	}

	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
