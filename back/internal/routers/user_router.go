package routers

import (
	"back/global"
	"back/internal/controllers"
	"back/internal/dao"
	"back/internal/services"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitApiRouter(parent *gin.RouterGroup) {
	publicRouter := parent.Group("users")

	//注入依赖
	UserCtrl := controllers.NewUserController(
		services.NewUserService(
			dao.NewUserDAO(global.DB)))
	{
		// 用户注册
		publicRouter.POST("/register", UserCtrl.Register)

		// 用户登录
		publicRouter.POST("/login", UserCtrl.Login)
	}
}