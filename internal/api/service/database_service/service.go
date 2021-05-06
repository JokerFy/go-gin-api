package database_service

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/database_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
)

var _ Service = (*service)(nil)

// 定义缓存前缀
var cacheKeyPrefix = configs.ProjectName() + ":authorized:"

type Service interface {
	i()
	List(ctx core.Context, searchData *SearchData) (listData []*database_repo.Database, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*database_repo.Database, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*database_repo.Database, err error)
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
