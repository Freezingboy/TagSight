package routers

import (
	"back/global"
	"back/internal/controllers"
	"back/internal/dao"
	"back/internal/services"
	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (fr *FileRouter) InitApiRouter(parent *gin.RouterGroup) {
	privateRouter := parent.Group("files")

	//TODO后续可以添加权限控制

	//注入依赖
	FileCtrl := controllers.NewFileController(
		services.NewFileService(
			dao.NewFileDao(global.DB)))
	{
		privateRouter.POST("upload", FileCtrl.UploadFile)
	}

}
