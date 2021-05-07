package template_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/database_service"
	"github.com/xinliangnote/go-gin-api/internal/api/service/template_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	// Dbs 查询 DB
	// @Tags API.tool
	// @Router /api/tool/data/dbs [get]
	Dbs() core.HandlerFunc

	// Tables 查询 Table
	// @Tags API.tool
	// @Router /api/tool/data/tables [post]
	Tables() core.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	db              db.Repo
	cache           cache.Repo
	hashids         hash.Hash
	databaseService database_service.Service
	templateService template_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:          logger,
		db:              db,
		cache:           cache,
		hashids:         hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		databaseService: database_service.New(db, cache),
		templateService: template_service.New(db, cache),
	}
}

func (h *handler) i() {}
