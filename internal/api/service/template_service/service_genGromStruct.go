package template_service

import (
	"bytes"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/database_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/template_repo"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"strings"
)

func (s *service) GenGormStruct(ctx core.Context, dbInfo *database_repo.Database, tableName string) (res string, err error) {
	qb := template_repo.NewQueryBuilder()
	qb.WhereId(db_repo.EqualPredicate, 1)
	/*	info, err := qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
		if err != nil {
			return "", err
		}*/
	//res = info.ModelTemplate
	res = genGormStruct(tableName, dbInfo)
	return res, nil
}

//生成gorm结构体
func genGormStruct(tableName string, dbInfo *database_repo.Database) string {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbInfo.Account, dbInfo.Password, dbInfo.Addr+":"+fmt.Sprintf("%v", dbInfo.Ports), "information_schema")
	db := sqlx.MustConnect("mysql", dns)
	defer db.Close()
	type FieldInfo struct {
		ColName    string `db:"COLUMN_NAME"`
		DataType   string `db:"DATA_TYPE"`
		ColComment string `db:"COLUMN_COMMENT"`
		IsNullable string `db:"IS_NULLABLE"`
	}
	var fs []FieldInfo
	err := db.Select(&fs, "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT, IS_NULLABLE FROM COLUMNS WHERE TABLE_NAME=? and table_schema=?", tableName, dbInfo.Name)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var TplStructDefine string
	if len(fs) > 0 {
		var buffer bytes.Buffer
		buffer.WriteString("package models\n")
		buffer.WriteString("type " + fmtFieldDefine(tableName) + " struct {\n")
		for _, v := range fs {
			buffer.WriteString("" + fmtFieldDefine(v.ColName) + " ")
			switch v.DataType {
			case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
				buffer.WriteString("[]byte")

			case "bit", "int", "tinyint", "smallint", "mediumint":
				if strings.Contains(v.DataType, "unsigned") {
					buffer.WriteString("uint")
				} else {
					buffer.WriteString("int")
				}

			case "bigint":
				if strings.Contains(v.DataType, "unsigned") {
					buffer.WriteString("uint64")
				} else {
					buffer.WriteString("int64")
				}

			case "float", "double", "decimal":
				buffer.WriteString("float64")

			case "bool":
				buffer.WriteString("bool")

			case "datetime", "timestamp", "date", "time":
				buffer.WriteString("*time.Time")

			default:
				// Auto detecting type.
				switch {
				case strings.Contains(v.DataType, "int64"):
					buffer.WriteString("int64")
				case strings.Contains(v.DataType, "int"):
					buffer.WriteString("int")
				case strings.Contains(v.DataType, "text") || strings.Contains(v.DataType, "char"):
					buffer.WriteString("string")
				case strings.Contains(v.DataType, "float") || strings.Contains(v.DataType, "double"):
					buffer.WriteString("float64")
				case strings.Contains(v.DataType, "bool"):
					buffer.WriteString("bool")
				case strings.Contains(v.DataType, "binary") || strings.Contains(v.DataType, "blob"):
					buffer.WriteString("[]byte")
				case strings.Contains(v.DataType, "date") || strings.Contains(v.DataType, "time"):
					buffer.WriteString("*time.Time")
				default:
					buffer.WriteString("string")
				}
			}
			buffer.WriteString(fmt.Sprintf("		`gorm:\"%s\" json:\"%s\"`\n", v.ColName, v.ColName))
		}
		buffer.WriteString(`}`)
		TplStructDefine = buffer.String()
	}
	fmt.Sprintf("!232131222222222222222222222222222222")
	fmt.Sprintf(TplStructDefine)
	return TplStructDefine
}

//首字母大写
func strFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

//驼峰转换
func fmtFieldDefine(src string) string {
	temp := strings.Split(src, "_") // 有下划线的，需要拆分
	var str string
	for i := 0; i < len(temp); i++ {
		b := []rune(temp[i])
		for j := 0; j < len(b); j++ {
			if j == 0 {
				// 首字母大写转换
				b[j] -= 32
				str += string(b[j])
			} else {
				str += string(b[j])
			}
		}
	}
	return str
}
