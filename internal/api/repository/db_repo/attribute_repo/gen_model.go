package attribute_repo

import "time"

//
//go:generate gormgen -structs Attribute -input .
type Attribute struct {
	Id        int64     //
	SpuId     string    // 主题编号
	OptName   string    // 标题
	Img       string    // 详情图
	Value     string    // 值
	Status    int32     // 是否禁用   1启用   非1禁用
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
}
