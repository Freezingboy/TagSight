package router

import (
	"back/internal/routers"
	"github.com/gin-gonic/gin"
)

type AllRouter struct {
	FileRouter routers.FileRouter
}

func RouterInit() *gin.Engine {
	r := gin.Default()
	allRouter := new(AllRouter)
	v1 := r.Group("/v1")
	{
		allRouter.FileRouter.InitApiRouter(v1)
	}

	return r
}
