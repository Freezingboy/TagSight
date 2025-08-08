package routers

import (
	"back/global"
	"back/internal/controllers"
	"back/internal/dao"
	"back/internal/services"
	"github.com/gin-gonic/gin"
)

type FileTagRouter struct{}

func (ftr *FileTagRouter) InitApiRouter(parent *gin.RouterGroup) {
	privateRouter := parent.Group("")

	//注入依赖
	FileTagCtrl := controllers.NewFileTagController(
		services.NewFileTagService(
			dao.NewFileTagRelationDAO(global.DB),
			dao.NewFileDAO(global.DB),
			dao.NewTagDAO(global.DB)))
	{
		// 添加标签到文件
		privateRouter.POST("/file-tags", FileTagCtrl.AddTagToFile)

		// 从文件移除标签
		privateRouter.DELETE("/file-tags", FileTagCtrl.RemoveTagFromFile)

		// 获取文件标签
		privateRouter.GET("/files/:fileId/tags", FileTagCtrl.GetFileTags)
	}
}