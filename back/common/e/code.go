package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"back/common"
)

const (
	SUCCESS                  = 1    // 成功
	ERROR                    = 2    // 内部错误
	UNKNOW_IDENTITY          = 403  // 未知身份
	MysqlERR                 = 1001 // MySQL出错
	MysqlTransActionERR      = 1002 // MySQL事务执行出错
	RedisERR                 = 1003 // Redis出错
	
	// 用户相关错误码 (2000-2099)
	ErrorPasswordError       = 2001 // 密码错误
	ErrorAccountNotFound     = 2002 // 账号不存在
	ErrorAccountLocked       = 2003 // 账号被锁定
	ErrorUserNotLogin        = 2004 // 用户未登录
	ErrorLoginFailed         = 2005 // 登录失败
	ErrorPasswordEditFailed  = 2006 // 密码修改失败
	ErrorUserExists          = 2007 // 用户名已存在
	ErrorInvalidCredentials  = 2008 // 无效的凭证
	
	// 标签相关错误码 (2100-2199)
	ErrorTagExists           = 2100 // 标签已存在
	ErrorTagNotFound         = 2101 // 标签不存在
	ErrorTagNameInvalid      = 2102 // 标签名称无效
	ErrorTagColorInvalid     = 2103 // 标签颜色无效
	ErrorTagCreateFailed     = 2104 // 标签创建失败
	ErrorTagUpdateFailed     = 2105 // 标签更新失败
	ErrorTagDeleteFailed     = 2106 // 标签删除失败
	ErrorTagBeRelatedByFile  = 2107 // 标签被文件关联，不能删除
	
	// 文件相关错误码 (2200-2299)
	ErrorFileNotFound        = 2200 // 文件不存在
	ErrorFileUploadFailed    = 2201 // 文件上传失败
	ErrorFileDownloadFailed  = 2202 // 文件下载失败
	ErrorFileDeleteFailed    = 2203 // 文件删除失败
	ErrorFileTypeNotAllowed  = 2204 // 文件类型不允许
	ErrorFileSizeTooLarge    = 2205 // 文件大小超过限制
	ErrorFileNameInvalid     = 2206 // 文件名无效
	ErrorFilePathInvalid     = 2207 // 文件路径无效
	
	// 文件标签关联相关错误码 (2300-2399)
	ErrorRelationExists      = 2300 // 关联已存在
	ErrorRelationNotFound    = 2301 // 关联不存在
	ErrorRelationCreateFailed = 2302 // 关联创建失败
	ErrorRelationDeleteFailed = 2303 // 关联删除失败
	
	// 其他错误码
	ErrorUnknowError         = 9000 // 未知错误
	ErrorInvalidParameter    = 9001 // 无效的参数
	ErrorNotAuthorized       = 9002 // 没有权限
	ErrorOperationFailed     = 9003 // 操作失败
)

var ErrMsg = map[int]string{
	SUCCESS:                  "成功",
	ERROR:                    "内部错误",
	UNKNOW_IDENTITY:          "未知身份",
	MysqlERR:                 "数据库错误",
	MysqlTransActionERR:      "数据库事务执行错误",
	RedisERR:                 "缓存服务错误",
	
	// 用户相关错误信息
	ErrorPasswordError:       "密码错误",
	ErrorAccountNotFound:     "账号不存在",
	ErrorAccountLocked:       "账号已被锁定",
	ErrorUserNotLogin:        "用户未登录",
	ErrorLoginFailed:         "登录失败",
	ErrorPasswordEditFailed:  "密码修改失败",
	ErrorUserExists:          "用户名已存在",
	ErrorInvalidCredentials:  "无效的凭证",
	
	// 标签相关错误信息
	ErrorTagExists:           "标签已存在",
	ErrorTagNotFound:         "标签不存在",
	ErrorTagNameInvalid:      "标签名称无效",
	ErrorTagColorInvalid:     "标签颜色无效",
	ErrorTagCreateFailed:     "标签创建失败",
	ErrorTagUpdateFailed:     "标签更新失败",
	ErrorTagDeleteFailed:     "标签删除失败",
	ErrorTagBeRelatedByFile:  "标签被文件关联，不能删除",
	
	// 文件相关错误信息
	ErrorFileNotFound:        "文件不存在",
	ErrorFileUploadFailed:    "文件上传失败",
	ErrorFileDownloadFailed:  "文件下载失败",
	ErrorFileDeleteFailed:    "文件删除失败",
	ErrorFileTypeNotAllowed:  "文件类型不允许",
	ErrorFileSizeTooLarge:    "文件大小超过限制",
	ErrorFileNameInvalid:     "文件名无效",
	ErrorFilePathInvalid:     "文件路径无效",
	
	// 文件标签关联相关错误信息
	ErrorRelationExists:      "关联已存在",
	ErrorRelationNotFound:    "关联不存在",
	ErrorRelationCreateFailed: "关联创建失败",
	ErrorRelationDeleteFailed: "关联删除失败",
	
	// 其他错误信息
	ErrorUnknowError:         "未知错误",
	ErrorInvalidParameter:    "无效的参数",
	ErrorNotAuthorized:       "没有权限",
	ErrorOperationFailed:     "操作失败",
}

func GetMsg(code int) string {
	msg, ok := ErrMsg[code]
	if ok {
		return msg
	}
	return "未知错误码"
}

func Send(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: GetMsg(code)})
}

func SendWithData(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: GetMsg(code), Data: data})
}