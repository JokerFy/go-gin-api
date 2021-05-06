package template_handler

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/xinliangnote/go-gin-api/cmd/mysqlmd/mysql"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/api/service/database_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type tablesRequest struct {
	DbName string `form:"db_name"` // 数据库名称
	Gen    bool   `form:"gen"`
}

type tablesResponse struct {
	List     []tableData `json:"list"`      // 数据表列表
	GormData string      `json:"gorm_data"` // 数据表列表
}

type tableData struct {
	Name    string `json:"table_name"`    // 数据表名称
	Comment string `json:"table_comment"` // 数据表备注
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
func (h *handler) Tables() core.HandlerFunc {
	return func(c core.Context) {
		req := new(tablesRequest)
		res := new(tablesResponse)
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
			// 初始化 DB
			tableDb, _ := mysql.New(databaseList[0].Addr, databaseList[0].Account, databaseList[0].Password, databaseList[0].Name)
			tables, _ := QueryTables(tableDb.GetDb(), databaseList[0].Name, "")
			for _, v := range tables {
				var info tableData
				info.Name = v.Name
				info.Comment = ""
				res.List = append(res.List, info)
			}

			if req.Gen {
				shellPath := fmt.Sprintf("./scripts/gormgen.sh %s %s %s %s %s", databaseList[0].Addr, databaseList[0].Account, databaseList[0].Password, databaseList[0].Name, "attribute")

				// runtime.GOOS = linux or darwin
				command := exec.Command("/bin/bash", "-c", shellPath)
				if runtime.GOOS == "windows" {
					command = exec.Command("cmd", "/C", shellPath)
				}

				var stderr bytes.Buffer
				command.Stderr = &stderr

				output, err := command.Output()
				fmt.Print(output)
				if err != nil {
					c.Payload(stderr.String())
					return
				}
			}

			c.Payload(res)
		}

	}
}
