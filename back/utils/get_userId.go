package utils

import (
	"back/common/enum"
	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(ctx *gin.Context) uint64 {
	userID, exists := ctx.Get(enum.CurrentId)
	if !exists {
		return 0
	}
	return userID.(uint64)
}
