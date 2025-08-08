package routers

import (
	"back/global"
	"back/internal/controllers"
	"back/internal/dao"
	"back/internal/services"
	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (tr *TagRouter) InitApiRouter(parent *gin.RouterGroup) {
	privateRouter := parent.Group("tags")

	//注入依赖
	TagCtrl := controllers.NewTagController(
		services.NewTagService(
			dao.NewTagDAO(global.DB),
			dao.NewFileTagRelationDAO(global.DB)))
	{
		// 创建标签
		privateRouter.POST("", TagCtrl.CreateTag)

		// 获取用户标签
		privateRouter.GET("", TagCtrl.GetUserTags)

		// 更新标签
		privateRouter.PUT("/:tagId", TagCtrl.UpdateTag)

		// 删除标签
		privateRouter.DELETE("/:tagId", TagCtrl.DeleteTag)
	}
}
