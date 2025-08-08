package request

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	Name  string `json:"name" binding:"required,max=50"`
	Color string `json:"color" binding:"required,max=20"`
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	Name  string `json:"name" binding:"required,max=50"`
	Color string `json:"color" binding:"required,max=20"`
}
