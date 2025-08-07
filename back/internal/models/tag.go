package models

// Tag 标签表实体类
type Tag struct {
	ID     uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:标签ID"`
	UserID uint64 `gorm:"column:user_id;type:bigint unsigned;not null;comment:所属用户ID"`
	Name   string `gorm:"column:name;type:varchar(50);not null;comment:标签名"`
	Color  string `gorm:"column:color;type:varchar(20);comment:标签颜色（如 #FF0000）"`
}

// TableName 指定Tag结构体对应的数据库表名
func (t *Tag) TableName() string {
	return "tag"
}
