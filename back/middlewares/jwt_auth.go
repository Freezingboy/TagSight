package middlewares

import (
	"back/common"
	"back/common/e"
	"back/common/enum"
	"back/common/utils"
	"back/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.Request.Header.Get(config.Jwt.Name)
		// 解析获取用户载荷信息
		payLoad, err := utils.ParseToken(token, config.Jwt.Secret)
		if err != nil {
			code = e.UNKNOW_IDENTITY
			c.JSON(http.StatusUnauthorized, common.Response{
				Code:    code,
				Message: e.GetMsg(code),
			})
			c.Abort()
			return
		}
		// 在上下文设置载荷信息
		c.Set(enum.CurrentId, payLoad.UserId)
		c.Set(enum.CurrentName, payLoad.GrantScope)
		// 这里是否要通知客户端重新保存新的Token
		c.Next()
	}
}
