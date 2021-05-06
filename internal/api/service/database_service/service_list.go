package database_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/database_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

func (s *service) List(ctx core.Context, searchData *SearchData) (listData []*database_repo.Database, err error) {

	qb := database_repo.NewQueryBuilder()
	if searchData.DataBaseName != "" {
		qb.WhereName(db_repo.EqualPredicate, searchData.DataBaseName)
	}

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
