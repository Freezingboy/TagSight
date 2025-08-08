package router

import (
	"back/internal/routers"
	"github.com/gin-gonic/gin"
)

type AllRouter struct {
	routers.FileRouter
	routers.TagRouter
	routers.UserRouter
	routers.FileTagRouter
}

func RouterInit() *gin.Engine {
	r := gin.Default()

	// 允许跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	allRouter := new(AllRouter)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	// 公开路由
	{
		allRouter.UserRouter.InitApiRouter(v1)
		allRouter.FileTagRouter.InitApiRouter(v1)
		allRouter.TagRouter.InitApiRouter(v1)
		allRouter.FileRouter.InitApiRouter(v1)
	}
	return r
}
