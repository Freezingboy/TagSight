package controllers

import (
	"back/api/request"
	"back/common"
	"net/http"

	"github.com/gin-gonic/gin"

	"back/internal/services"
)

// UserController 用户控制器
type UserController struct {
	userService *services.UserService
}

// NewUserController 创建UserController实例
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 创建新用户账号
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body services.RegisterRequest true "注册信息"
// @Success 200 {object} Response{data=services.RegisterResponse} "注册成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 409 {object} Response "用户名已存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /users/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req services.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	resp, err := c.userService.Register(&req)
	if err != nil {
		if err == services.ErrUserExists {
			common.ResponseError(ctx, http.StatusConflict, "用户名已存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "注册失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "注册成功", resp)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录并获取令牌
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body services.LoginRequest true "登录信息"
// @Success 200 {object} Response{data=services.LoginResponse} "登录成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "用户名或密码错误"
// @Failure 500 {object} Response "服务器错误"
// @Router /users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	resp, err := c.userService.Login(&req)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			common.ResponseError(ctx, http.StatusUnauthorized, "用户名或密码错误")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "登录失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "登录成功", resp)
}
