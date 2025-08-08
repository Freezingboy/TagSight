package dao

import (
	"errors"

	"gorm.io/gorm"

	"back/internal/models"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	DB *gorm.DB
}

// NewUserDAO 创建UserDAO实例
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{DB: db}
}

// Create 创建用户
func (dao *UserDAO) Create(user *models.User) error {
	return dao.DB.Create(user).Error
}

// GetByUsername 根据用户名查询用户
func (dao *UserDAO) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := dao.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在返回nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByID 根据ID查询用户
func (dao *UserDAO) GetByID(id uint64) (*models.User, error) {
	var user models.User
	err := dao.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (dao *UserDAO) Update(user *models.User) error {
	return dao.DB.Save(user).Error
}

// Delete 删除用户
func (dao *UserDAO) Delete(id uint64) error {
	return dao.DB.Delete(&models.User{}, id).Error
}
