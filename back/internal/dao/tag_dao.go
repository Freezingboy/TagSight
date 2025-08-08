package dao

import (
	"errors"

	"gorm.io/gorm"

	"back/internal/models"
)

// TagDAO 标签数据访问对象
type TagDAO struct {
	DB *gorm.DB
}

// NewTagDAO 创建TagDAO实例
func NewTagDAO(db *gorm.DB) *TagDAO {
	return &TagDAO{DB: db}
}

// Create 创建标签
func (dao *TagDAO) Create(tag *models.Tag) error {
	return dao.DB.Create(tag).Error
}

// GetByID 根据ID查询标签
func (dao *TagDAO) GetByID(id uint64, userID uint64) (*models.Tag, error) {
	var tag models.Tag
	err := dao.DB.Where("id = ? AND user_id = ?", id, userID).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

// GetByName 根据名称查询标签
func (dao *TagDAO) GetByName(name string, userID uint64) (*models.Tag, error) {
	var tag models.Tag
	err := dao.DB.Where("name = ? AND user_id = ?", name, userID).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

// ListByUserID 查询用户的所有标签
func (dao *TagDAO) ListByUserID(userID uint64) ([]models.Tag, error) {
	var tags []models.Tag
	err := dao.DB.Where("user_id = ?", userID).Find(&tags).Error
	return tags, err
}

// Update 更新标签
func (dao *TagDAO) Update(tag *models.Tag) error {
	return dao.DB.Save(tag).Error
}

// Delete 删除标签
func (dao *TagDAO) Delete(id uint64, userID uint64) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Tag{}).Error
}
