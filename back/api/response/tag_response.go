package response

// TagResponse 标签响应
type TagResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// CreateTagResponse 创建标签响应
type CreateTagResponse struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"userId"`
	Name   string `json:"name"`
	Color  string `json:"color"`
}
