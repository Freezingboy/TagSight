package gorm

import (
	"back/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 数据库初始化
func InitDataBase(mysqlCfg *config.MysqlSetting) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlCfg.Dsn(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	return db
}
