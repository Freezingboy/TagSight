package models

// User 用户表实体类
type User struct {
	ID       uint64 `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:用户ID"`
	Username string `gorm:"column:username;type:varchar(50);not null;uniqueIndex:idx_username;comment:用户名"`
	Password string `gorm:"column:password;type:varchar(255);not null;comment:密码（加密存储）"`
}

// TableName 指定User结构体对应的数据库表名
func (u *User) TableName() string {
	return "user"
}
