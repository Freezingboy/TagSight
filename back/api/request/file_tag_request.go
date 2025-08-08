package request

// FileTagRequest 文件标签请求
type FileTagRequest struct {
	FileID uint64 `json:"fileId" binding:"required"`
	TagID  uint64 `json:"tagId" binding:"required"`
}