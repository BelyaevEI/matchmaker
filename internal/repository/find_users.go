package repository

import (
	"context"
	"sort"
	"strconv"

	"github.com/BelyaevEI/matchmaker/internal/model"
	"github.com/BelyaevEI/matchmaker/internal/utils"
	"github.com/BelyaevEI/platform_common/pkg/db"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) FindUsersForMatch(ctx context.Context, userOld model.User) ([]model.User, error) {
	if r.storageFlag {
		return r.findUsersMem(ctx, userOld)
	}

	return r.findUsersDB(ctx, userOld)
}

func (r *repository) findUsersMem(_ context.Context, userOld model.User) ([]model.User, error) {

	if len(r.poolUsers) < int(r.groupsize) {
		return nil, nil
	}

	sort.Slice(r.poolUsers, func(i, j int) bool {
		return utils.DistanceMin(r.poolUsers[i], userOld) < utils.DistanceMin(r.poolUsers[j], userOld)
	})

	return r.poolUsers[:r.groupsize+1], nil
}

func (r *repository) findUsersDB(ctx context.Context, userOld model.User) ([]model.User, error) {

	users := make([]model.User, 0)
	skill := strconv.Itoa(int(userOld.Skill))
	latency := strconv.Itoa(int(userOld.Latency))

	builder := sq.Select(nameColumn, skillColumn, latencyColumn, createdAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		OrderBy("POWER(skill - ?, 2) + POWER(latency - ?, 2) ASC",
			skill, latency).Limit(uint64(r.groupsize))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "find users",
		QueryRaw: query,
	}

	err = r.db.DB().ScanAllContext(ctx, &users, q, args...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
