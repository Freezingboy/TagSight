package response

import (
	"time"
)

// FileResponse 文件响应
type FileResponse struct {
	ID         uint64        `json:"id"`
	Name       string        `json:"name"`
	Size       int           `json:"size"`
	Type       string        `json:"type"`
	Path       string        `json:"path,omitempty"` // 内部使用，API响应中省略
	UploadTime time.Time     `json:"uploadTime"`
	Tags       []TagResponse `json:"tags,omitempty"` // 可选字段，查询文件列表时包含
}

// FileListResponse 文件列表响应
type FileListResponse struct {
	Total int64          `json:"total"`
	List  []FileResponse `json:"list"`
}
