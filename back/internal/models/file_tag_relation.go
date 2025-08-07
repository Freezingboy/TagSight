package models

// FileTagRelation 文件-标签关联表实体类
type FileTagRelation struct {
	ID     uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:关联ID"`
	UserID uint64 `gorm:"column:user_id;type:bigint unsigned;not null;comment:所属用户ID"`
	FileID uint64 `gorm:"column:file_id;type:bigint unsigned;not null;index:idx_file;comment:文件ID"`
	TagID  uint64 `gorm:"column:tag_id;type:bigint unsigned;not null;index:idx_tag;comment:标签ID"`
}

// TableName 指定FileTagRelation结构体对应的数据库表名
func (ftr *FileTagRelation) TableName() string {
	return "file_tag_relation"
}
