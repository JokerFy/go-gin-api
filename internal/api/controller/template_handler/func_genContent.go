package template_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/database_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"net/http"
)

type genContentRequest struct {
	DbName    string `form:"db_name"`    // 数据库名称
	TableName string `form:"table_name"` // 数据库名称
}

type genContentResponse struct {
	GormData string `json:"gorm_data"` // 数据表列表
}

// Tables 查询 Table
// @Summary 查询 Table
// @Description 查询 Table
// @Tags API.tool
// @Accept multipart/form-data
// @Produce json
// @Param db_name formData string true "数据库名称"
// @Success 200 {object} tablesResponse
// @Failure 400 {object} code.Failure
// @Router /api/tool/data/tables [post]
func (h *handler) GenContent() core.HandlerFunc {
	return func(c core.Context) {
		req := new(genContentRequest)
		res := new(genContentResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		searchOneData := new(database_service.SearchData)
		searchOneData.DataBaseName = req.DbName
		databaseList, _ := h.databaseService.List(c, searchOneData)
		if len(databaseList) < 1 {

		} else {
			templateInfo, _ := h.templateService.GenGormStruct(c, databaseList[0], req.TableName)
			res.GormData = templateInfo
		}
		c.Payload(res)
	}
}
