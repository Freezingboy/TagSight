package dao

import (
	"errors"

	"gorm.io/gorm"

	"back/internal/models"
)

// FileDAO 文件数据访问对象
type FileDAO struct {
	DB *gorm.DB
}

// NewFileDAO 创建FileDAO实例
func NewFileDAO(db *gorm.DB) *FileDAO {
	return &FileDAO{DB: db}
}

// Create 创建文件记录
func (dao *FileDAO) Create(file *models.File) error {
	return dao.DB.Create(file).Error
}

// GetByID 根据ID查询文件
func (dao *FileDAO) GetByID(id uint64, userID uint64) (*models.File, error) {
	var file models.File
	err := dao.DB.Where("id = ? AND user_id = ?", id, userID).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

// GetByPath 根据路径查询文件
func (dao *FileDAO) GetByPath(path string) (*models.File, error) {
	var file models.File
	err := dao.DB.Where("path = ?", path).First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &file, nil
}

// ListByUserID 分页查询用户的文件
func (dao *FileDAO) ListByUserID(userID uint64, page, size int) ([]models.File, int64, error) {
	var files []models.File
	var total int64

	// 查询总数
	err := dao.DB.Model(&models.File{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * size
	err = dao.DB.Where("user_id = ?", userID).
		Offset(offset).
		Limit(size).
		Order("upload_time DESC").
		Find(&files).Error

	return files, total, err
}

// ListByTagID 根据标签ID查询文件
func (dao *FileDAO) ListByTagID(userID uint64, tagID uint64, page, size int) ([]models.File, int64, error) {
	var files []models.File
	var total int64

	// 查询总数
	err := dao.DB.Model(&models.File{}).
		Joins("JOIN file_tag_relation ON file.id = file_tag_relation.file_id").
		Where("file.user_id = ? AND file_tag_relation.tag_id = ?", userID, tagID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * size
	err = dao.DB.
		Joins("JOIN file_tag_relation ON file.id = file_tag_relation.file_id").
		Where("file.user_id = ? AND file_tag_relation.tag_id = ?", userID, tagID).
		Offset(offset).
		Limit(size).
		Order("file.upload_time DESC").
		Find(&files).Error

	return files, total, err
}

// Update 更新文件信息
func (dao *FileDAO) Update(file *models.File) error {
	return dao.DB.Save(file).Error
}

// Delete 删除文件
func (dao *FileDAO) Delete(id uint64, userID uint64) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.File{}).Error
}
