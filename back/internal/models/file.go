package models

import "time"

// File 文件表实体类
type File struct {
	ID         uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:文件ID"`
	UserID     uint64    `gorm:"column:user_id;type:bigint unsigned;not null;index:idx_user;comment:所属用户ID"`
	Name       string    `gorm:"column:name;type:varchar(255);not null;comment:文件名"`
	Size       int       `gorm:"column:size;type:int;not null;comment:文件大小（字节）"`
	Type       string    `gorm:"column:type;type:varchar(50);not null;comment:文件类型（如 pdf、jpg）"`
	Path       string    `gorm:"column:path;type:varchar(255);not null;uniqueIndex:idx_path;comment:文件存储路径"`
	UploadTime time.Time `gorm:"column:upload_time;type:datetime;not null;default:current_timestamp;comment:上传时间"`
}

// TableName 指定File结构体对应的数据库表名
func (f *File) TableName() string {
	return "file"
}
