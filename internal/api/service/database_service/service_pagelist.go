package database_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/database_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchData struct {
	Page         int    `json:"page"`      // 第几页
	PageSize     int    `json:"page_size"` // 每页显示条数
	DataBaseName string `json:"name"`
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*database_repo.Database, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	name := searchData.DataBaseName
	if name != "" {
		name = searchData.DataBaseName
	}

	offset := (page - 1) * pageSize

	qb := database_repo.NewQueryBuilder()
	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
