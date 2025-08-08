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

	//注入依赖
	FileCtrl := controllers.NewFileController(
		services.NewFileService(
			dao.NewFileDAO(global.DB),
			dao.NewFileTagRelationDAO(global.DB),
			dao.NewTagDAO(global.DB),
			"static"))
	{
		// 文件上传
		privateRouter.POST("/upload", FileCtrl.UploadFile)

		// 获取文件列表
		privateRouter.GET("", FileCtrl.GetFileList)

		// 下载文件
		privateRouter.GET("/:fileId/download", FileCtrl.DownloadFile)

		// 删除文件
		privateRouter.DELETE("/:fileId", FileCtrl.DeleteFile)
	}
}
