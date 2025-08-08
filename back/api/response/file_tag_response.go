package response

// FileTagResponse 文件标签响应
type FileTagResponse struct {
	ID     uint64 `json:"id"`
	FileID uint64 `json:"fileId"`
	TagID  uint64 `json:"tagId"`
}