package services

import (
	"back/api/request"
	"back/api/response"
	"back/common/utils"
	"back/config"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"back/internal/dao"
	"back/internal/models"
)

var (
	// ErrUserExists 用户已存在错误
	ErrUserExists = errors.New("用户名已存在")
	// ErrInvalidCredentials 无效的凭证错误
	ErrInvalidCredentials = errors.New("用户名或密码错误")
)

// UserService 用户服务
type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建UserService实例
func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=1,max=50"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterResponse 用户注册响应
type RegisterResponse struct {
	ID         uint64    `json:"id"`
	Username   string    `json:"username"`
	CreateTime time.Time `json:"createTime"`
}

// Register 用户注册
func (s *UserService) Register(req *RegisterRequest) (*RegisterResponse, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userDAO.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserExists
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err := s.userDAO.Create(user); err != nil {
		return nil, err
	}

	// 返回响应
	return &RegisterResponse{
		ID:         user.ID,
		Username:   user.Username,
		CreateTime: time.Now(),
	}, nil
}

// Login 用户登录
func (s *UserService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	// 查询用户
	user, err := s.userDAO.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, config.Jwt.Name, config.Jwt.Secret)

	if err != nil {
		return nil, err
	}

	// 返回响应
	resp := &response.LoginResponse{
		Token: token,
	}
	resp.User.ID = user.ID
	resp.User.Username = user.Username

	return resp, nil
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(userID uint64) (*models.User, error) {
	return s.userDAO.GetByID(userID)
}
