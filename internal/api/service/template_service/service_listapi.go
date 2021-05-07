package template_service

import (
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/template_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type SearchAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
}

func (s *service) ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*template_repo.Template, err error) {

	qb := template_repo.NewQueryBuilder()

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
