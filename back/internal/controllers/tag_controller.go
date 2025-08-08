package controllers

import (
	"back/api/request"
	"back/common"
	"back/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"back/internal/services"
)

// TagController 标签控制器
type TagController struct {
	tagService *services.TagService
}

// NewTagController 创建TagController实例
func NewTagController(tagService *services.TagService) *TagController {
	return &TagController{
		tagService: tagService,
	}
}

// CreateTag 创建标签
// @Summary 创建标签
// @Description 创建新标签
// @Tags 标签管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body services.CreateTagRequest true "标签信息"
// @Success 200 {object} Response{data=services.CreateTagResponse} "标签创建成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 409 {object} Response "标签名已存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /tags [post]
func (c *TagController) CreateTag(ctx *gin.Context) {
	var req request.CreateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	resp, err := c.tagService.CreateTag(&req, userID)
	if err != nil {
		if err == services.ErrTagExists {
			common.ResponseError(ctx, http.StatusConflict, "标签名已存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "创建标签失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "标签创建成功", resp)
}

// GetUserTags 获取用户标签列表
// @Summary 获取用户标签列表
// @Description 获取当前用户的所有标签
// @Tags 标签管理
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Response{data=[]services.TagResponse} "查询成功"
// @Failure 401 {object} Response "未授权"
// @Failure 500 {object} Response "服务器错误"
// @Router /tags [get]
func (c *TagController) GetUserTags(ctx *gin.Context) {
	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	tags, err := c.tagService.GetUserTags(userID)
	if err != nil {
		common.ResponseError(ctx, http.StatusInternalServerError, "查询标签失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "查询成功", tags)
}

// UpdateTag 更新标签
// @Summary 更新标签
// @Description 更新标签信息
// @Tags 标签管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param tagId path int true "标签ID"
// @Param request body services.UpdateTagRequest true "标签信息"
// @Success 200 {object} Response{data=services.TagResponse} "标签更新成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "标签不存在"
// @Failure 409 {object} Response "标签名已存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /tags/{tagId} [put]
func (c *TagController) UpdateTag(ctx *gin.Context) {
	// 解析路径参数
	tagIDStr := ctx.Param("tagId")
	tagID, err := strconv.ParseUint(tagIDStr, 10, 64)
	if err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "无效的标签ID")
		return
	}

	var req services.UpdateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	resp, err := c.tagService.UpdateTag(tagID, &req, userID)
	if err != nil {
		switch err {
		case services.ErrTagNotFound:
			common.ResponseError(ctx, http.StatusNotFound, "标签不存在")
		case services.ErrTagExists:
			common.ResponseError(ctx, http.StatusConflict, "标签名已存在")
		default:
			common.ResponseError(ctx, http.StatusInternalServerError, "更新标签失败: "+err.Error())
		}
		return
	}

	common.ResponseSuccess(ctx, "标签更新成功", resp)
}

// DeleteTag 删除标签
// @Summary 删除标签
// @Description 删除标签及其关联
// @Tags 标签管理
// @Produce json
// @Security BearerAuth
// @Param tagId path int true "标签ID"
// @Success 200 {object} Response "标签删除成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "标签不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /tags/{tagId} [delete]
func (c *TagController) DeleteTag(ctx *gin.Context) {
	// 解析路径参数
	tagIDStr := ctx.Param("tagId")
	tagID, err := strconv.ParseUint(tagIDStr, 10, 64)
	if err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "无效的标签ID")
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	err = c.tagService.DeleteTag(tagID, userID)
	if err != nil {
		if err == services.ErrTagNotFound {
			common.ResponseError(ctx, http.StatusNotFound, "标签不存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "删除标签失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "标签删除成功", nil)
}
