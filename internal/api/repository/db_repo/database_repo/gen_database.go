///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package database_repo

import (
	"fmt"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Database {
	return new(Database)
}

func NewQueryBuilder() *databaseRepoQueryBuilder {
	return new(databaseRepoQueryBuilder)
}

func (t *Database) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *Database) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (t *Database) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(&Database{}).Where("id = ?", t.Id).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

type databaseRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *databaseRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *databaseRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Database{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *databaseRepoQueryBuilder) First(db *gorm.DB) (*Database, error) {
	ret := &Database{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *databaseRepoQueryBuilder) QueryOne(db *gorm.DB) (*Database, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *databaseRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*Database, error) {
	var ret []*Database
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *databaseRepoQueryBuilder) Limit(limit int) *databaseRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *databaseRepoQueryBuilder) Offset(offset int) *databaseRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereIdIn(value []int32) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereIdNotIn(value []int32) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) OrderById(asc bool) *databaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereName(p db_repo.Predicate, value string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereNameIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereNameNotIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) OrderByName(asc bool) *databaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereAddr(p db_repo.Predicate, value string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "addr", p),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereAddrIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "addr", "IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereAddrNotIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "addr", "NOT IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) OrderByAddr(asc bool) *databaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "addr "+order)
	return qb
}

func (qb *databaseRepoQueryBuilder) WherePorts(p db_repo.Predicate, value int32) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ports", p),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WherePortsIn(value []int32) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ports", "IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WherePortsNotIn(value []int32) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ports", "NOT IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) OrderByPorts(asc bool) *databaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ports "+order)
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereAccount(p db_repo.Predicate, value string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "account", p),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereAccountIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "account", "IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WhereAccountNotIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "account", "NOT IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) OrderByAccount(asc bool) *databaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "account "+order)
	return qb
}

func (qb *databaseRepoQueryBuilder) WherePassword(p db_repo.Predicate, value string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", p),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WherePasswordIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) WherePasswordNotIn(value []string) *databaseRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "NOT IN"),
		value,
	})
	return qb
}

func (qb *databaseRepoQueryBuilder) OrderByPassword(asc bool) *databaseRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "password "+order)
	return qb
}
