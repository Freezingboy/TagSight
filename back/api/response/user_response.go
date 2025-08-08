package response

import "time"

// RegisterResponse 用户注册响应
type RegisterResponse struct {
	ID         uint64    `json:"id"`
	Username   string    `json:"username"`
	CreateTime time.Time `json:"createTime"`
}

// LoginResponse 用户登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}