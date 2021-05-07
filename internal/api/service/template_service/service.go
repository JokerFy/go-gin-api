package template_service

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/template_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

// 定义缓存前缀
var cacheKeyPrefix = configs.ProjectName() + ":authorized:"

type Service interface {
	i()
	Detail(ctx core.Context, id int32) (data *template_repo.Template, err error)
	List(ctx core.Context, searchData *SearchData) (listData []*template_repo.Template, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*template_repo.Template, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*template_repo.Template, err error)
	GenGormStruct(ctx core.Context) (res string, err error)
}

type service struct {
	db    db.Repo
	cache cache.Repo
}

func New(db db.Repo, cache cache.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
