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

// FileTagController 文件标签关联控制器
type FileTagController struct {
	fileTagService *services.FileTagService
}

// NewFileTagController 创建FileTagController实例
func NewFileTagController(fileTagService *services.FileTagService) *FileTagController {
	return &FileTagController{
		fileTagService: fileTagService,
	}
}

// AddTagToFile 为文件添加标签
// @Summary 为文件添加标签
// @Description 为指定文件添加标签
// @Tags 文件标签管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body services.FileTagRequest true "文件标签关联信息"
// @Success 200 {object} Response{data=services.FileTagResponse} "标签添加成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "文件或标签不存在"
// @Failure 409 {object} Response "文件已添加该标签"
// @Failure 500 {object} Response "服务器错误"
// @Router /file-tags [post]
func (c *FileTagController) AddTagToFile(ctx *gin.Context) {
	var req request.FileTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	resp, err := c.fileTagService.AddTagToFile(&req, userID)
	if err != nil {
		switch err {
		case services.ErrFileNotFound:
			common.ResponseError(ctx, http.StatusNotFound, "文件不存在")
		case services.ErrTagNotFound:
			common.ResponseError(ctx, http.StatusNotFound, "标签不存在")
		case services.ErrFileTagRelationExists:
			common.ResponseError(ctx, http.StatusConflict, "文件已添加该标签")
		default:
			common.ResponseError(ctx, http.StatusInternalServerError, "添加标签失败: "+err.Error())
		}
		return
	}

	common.ResponseSuccess(ctx, "标签添加成功", resp)
}

// RemoveTagFromFile 从文件移除标签
// @Summary 从文件移除标签
// @Description 从指定文件移除标签
// @Tags 文件标签管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body services.FileTagRequest true "文件标签关联信息"
// @Success 200 {object} Response "标签移除成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "文件标签关联不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /file-tags [delete]
func (c *FileTagController) RemoveTagFromFile(ctx *gin.Context) {
	var req request.FileTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	err := c.fileTagService.RemoveTagFromFile(&req, userID)
	if err != nil {
		if err == services.ErrFileTagRelationNotFound {
			common.ResponseError(ctx, http.StatusNotFound, "文件标签关联不存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "移除标签失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "标签移除成功", nil)
}

// GetFileTags 获取文件的标签列表
// @Summary 获取文件的标签列表
// @Description 获取指定文件的所有标签
// @Tags 文件标签管理
// @Produce json
// @Security BearerAuth
// @Param fileId path int true "文件ID"
// @Success 200 {object} Response{data=[]services.TagResponse} "查询成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "文件不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /files/{fileId}/tags [get]
func (c *FileTagController) GetFileTags(ctx *gin.Context) {
	// 解析路径参数
	fileIDStr := ctx.Param("fileId")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "无效的文件ID")
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	tags := c.fileTagService.GetFileTags(fileID, userID)
	if err != nil {
		if err == services.ErrFileNotFound {
			common.ResponseError(ctx, http.StatusNotFound, "文件不存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "查询标签失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "查询成功", tags)
}
