package controllers

import (
	"back/common"
	"back/utils"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"back/internal/services"
)

// FileController 文件控制器
type FileController struct {
	fileService *services.FileService
}

// NewFileController 创建FileController实例
func NewFileController(fileService *services.FileService) *FileController {
	return &FileController{
		fileService: fileService,
	}
}

// UploadFile 上传文件
// @Summary 上传文件
// @Description 上传文件并保存
// @Tags 文件管理
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "文件"
// @Success 200 {object} Response{data=services.FileResponse} "文件上传成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 409 {object} Response "文件路径已存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /files/upload [post]
func (c *FileController) UploadFile(ctx *gin.Context) {
	// 获取上传的文件
	form, err := ctx.MultipartForm()
	if err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "获取上传文件失败: "+err.Error())
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		common.ResponseError(ctx, http.StatusBadRequest, "请选择文件")
		return
	}
	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	// 上传文件
	resp, err := c.fileService.UploadFile(files, userID)
	if err != nil {
		if err == services.ErrFilePathExists {
			common.ResponseError(ctx, http.StatusConflict, "文件路径已存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "上传文件失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "文件上传成功", resp)
}

// GetFileList 获取文件列表
// @Summary 获取文件列表
// @Description 分页获取用户的文件列表，可按标签筛选
// @Tags 文件管理
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码，默认1"
// @Param size query int false "每页条数，默认20"
// @Param tagId query int false "标签ID，可选"
// @Success 200 {object} Response{data=services.FileListResponse} "查询成功"
// @Failure 401 {object} Response "未授权"
// @Failure 500 {object} Response "服务器错误"
// @Router /files [get]
func (c *FileController) GetFileList(ctx *gin.Context) {
	// 解析查询参数
	pageStr := ctx.DefaultQuery("page", "1")
	sizeStr := ctx.DefaultQuery("size", "20")
	tagIDStr := ctx.DefaultQuery("tagId", "0")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 || size > 100 {
		size = 20
	}

	tagID, err := strconv.ParseUint(tagIDStr, 10, 64)
	if err != nil {
		tagID = 0
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	// 查询文件列表
	resp, err := c.fileService.GetFileList(userID, page, size, tagID)
	if err != nil {
		common.ResponseError(ctx, http.StatusInternalServerError, "查询文件列表失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "查询成功", resp)
}

// DownloadFile 下载文件
// @Summary 下载文件
// @Description 下载指定ID的文件
// @Tags 文件管理
// @Produce octet-stream
// @Security BearerAuth
// @Param fileId path int true "文件ID"
// @Success 200 {file} binary "文件内容"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "文件不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /files/{fileId}/download [get]
func (c *FileController) DownloadFile(ctx *gin.Context) {
	// 解析路径参数
	fileIDStr := ctx.Param("fileId")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "无效的文件ID")
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	// 获取文件信息
	file, err := c.fileService.GetFile(fileID, userID)
	if err != nil {
		if err == services.ErrFileNotFound {
			common.ResponseError(ctx, http.StatusNotFound, "文件不存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "获取文件失败: "+err.Error())
		return
	}

	// 构建文件路径
	uploadDir := "." // 应从配置中获取
	filePath := filepath.Join(uploadDir, file.Path)

	// 设置响应头
	ctx.Header("Content-Disposition", "attachment; filename="+file.Name)
	ctx.Header("Content-Type", "application/octet-stream")

	// 返回文件
	ctx.File(filePath)
}

// DeleteFile 删除文件
// @Summary 删除文件
// @Description 删除指定ID的文件及其标签关联
// @Tags 文件管理
// @Produce json
// @Security BearerAuth
// @Param fileId path int true "文件ID"
// @Success 200 {object} Response "文件删除成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "文件不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /files/{fileId} [delete]
func (c *FileController) DeleteFile(ctx *gin.Context) {
	// 解析路径参数
	fileIDStr := ctx.Param("fileId")
	fileID, err := strconv.ParseUint(fileIDStr, 10, 64)
	if err != nil {
		common.ResponseError(ctx, http.StatusBadRequest, "无效的文件ID")
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserIDFromContext(ctx)

	// 删除文件
	err = c.fileService.DeleteFile(fileID, userID)
	if err != nil {
		if err == services.ErrFileNotFound {
			common.ResponseError(ctx, http.StatusNotFound, "文件不存在")
			return
		}
		common.ResponseError(ctx, http.StatusInternalServerError, "删除文件失败: "+err.Error())
		return
	}

	common.ResponseSuccess(ctx, "文件删除成功", nil)
}
