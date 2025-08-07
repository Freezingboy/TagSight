package initialize

import (
	"back/config"
	"back/global"
	"back/initialize/gorm"
	"back/initialize/router"
	"back/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func GlobalInit() *gin.Engine {
	configPath := pflag.StringP("config", "c", "config/app.yaml", "config file path")
	pflag.Parse()

	config.InitConfig(*configPath)      //初始化配置
	logger.InitLogger(config.LogOption) //初始化日志

	//数据库初始化
	global.DB = gorm.InitDataBase(config.MysqlOption)

	r := router.RouterInit()

	return r

}
