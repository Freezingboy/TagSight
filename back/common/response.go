package common

import "github.com/gin-gonic/gin"

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func ResponseSuccess(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, NewSuccessResponse(message, data))
}
func ResponseError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, NewErrorResponse(code, message))
}
