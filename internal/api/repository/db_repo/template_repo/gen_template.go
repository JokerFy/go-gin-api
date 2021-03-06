///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package template_repo

import (
	"fmt"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Template {
	return new(Template)
}

func NewQueryBuilder() *templateRepoQueryBuilder {
	return new(templateRepoQueryBuilder)
}

func (t *Template) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *Template) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (t *Template) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(&Template{}).Where("id = ?", t.Id).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

type templateRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *templateRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *templateRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Template{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *templateRepoQueryBuilder) First(db *gorm.DB) (*Template, error) {
	ret := &Template{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *templateRepoQueryBuilder) QueryOne(db *gorm.DB) (*Template, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *templateRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*Template, error) {
	var ret []*Template
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *templateRepoQueryBuilder) Limit(limit int) *templateRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *templateRepoQueryBuilder) Offset(offset int) *templateRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *templateRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereIdIn(value []int32) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereIdNotIn(value []int32) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderById(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WhereName(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereNameIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereNameNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByName(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WhereControllerTemplate(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "controller_template", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereControllerTemplateIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "controller_template", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereControllerTemplateNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "controller_template", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByControllerTemplate(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "controller_template "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WhereServiceTemplate(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "service_template", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereServiceTemplateIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "service_template", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereServiceTemplateNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "service_template", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByServiceTemplate(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "service_template "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WhereModelTemplate(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "model_template", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereModelTemplateIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "model_template", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereModelTemplateNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "model_template", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByModelTemplate(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "model_template "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WhereRouterTemplate(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "router_template", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereRouterTemplateIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "router_template", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereRouterTemplateNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "router_template", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByRouterTemplate(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "router_template "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WherePojoTemplate(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pojo_template", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WherePojoTemplateIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pojo_template", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WherePojoTemplateNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pojo_template", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByPojoTemplate(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pojo_template "+order)
	return qb
}

func (qb *templateRepoQueryBuilder) WhereProjectNamespace(p db_repo.Predicate, value string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "project_namespace", p),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereProjectNamespaceIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "project_namespace", "IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) WhereProjectNamespaceNotIn(value []string) *templateRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "project_namespace", "NOT IN"),
		value,
	})
	return qb
}

func (qb *templateRepoQueryBuilder) OrderByProjectNamespace(asc bool) *templateRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "project_namespace "+order)
	return qb
}
