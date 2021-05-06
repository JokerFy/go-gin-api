package template_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/database_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type dbsResponse struct {
	List []dbData `json:"list"` // 数据库列表
}

type dbData struct {
	DbName   string `json:"db_name"`  // 数据库名称
	Id       int32  `json:"id"`       //
	Addr     string `json:"addr"`     //
	Ports    int32  `json:"ports"`    //
	Account  string `json:"account"`  //
	Password string `json:"password"` //
}

// Dbs 查询 DB
// @Summary 查询 DB
// @Description 查询 DB
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} dbsResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/dbs [get]
func (h *handler) Dbs() core.HandlerFunc {
	return func(c core.Context) {
		res := new(dbsResponse)
		searchOneData := new(database_service.SearchData)
		databaseList, _ := h.databaseService.List(c, searchOneData)
		for _, v := range databaseList {
			data := dbData{
				DbName:   v.Name,
				Addr:     v.Addr,
				Id:       v.Id,
				Password: v.Password,
				Ports:    v.Ports,
				Account:  v.Account,
			}
			res.List = append(res.List, data)
		}

		c.Payload(res)
	}
}
