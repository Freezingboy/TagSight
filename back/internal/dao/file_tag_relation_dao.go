package dao

import (
	"errors"

	"gorm.io/gorm"

	"back/internal/models"
)

// FileTagRelationDAO 文件标签关联数据访问对象
type FileTagRelationDAO struct {
	DB *gorm.DB
}

// NewFileTagRelationDAO 创建FileTagRelationDAO实例
func NewFileTagRelationDAO(db *gorm.DB) *FileTagRelationDAO {
	return &FileTagRelationDAO{DB: db}
}

// Create 创建文件标签关联
func (dao *FileTagRelationDAO) Create(relation *models.FileTagRelation) error {
	return dao.DB.Create(relation).Error
}

// GetByFileAndTag 根据文件ID和标签ID查询关联
func (dao *FileTagRelationDAO) GetByFileAndTag(fileID, tagID, userID uint64) (*models.FileTagRelation, error) {
	var relation models.FileTagRelation
	err := dao.DB.Where("file_id = ? AND tag_id = ? AND user_id = ?", fileID, tagID, userID).First(&relation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &relation, nil
}

// ListTagsByFileID 查询文件的所有标签
func (dao *FileTagRelationDAO) ListTagsByFileID(fileID, userID uint64) ([]models.Tag, error) {
	var tags []models.Tag
	err := dao.DB.Table("tag").
		Joins("JOIN file_tag_relation ON tag.id = file_tag_relation.tag_id").
		Where("file_tag_relation.file_id = ? AND file_tag_relation.user_id = ?", fileID, userID).
		Find(&tags).Error
	return tags, err
}

// ListFilesByTagID 查询标签关联的所有文件
func (dao *FileTagRelationDAO) ListFilesByTagID(tagID, userID uint64) ([]models.File, error) {
	var files []models.File
	err := dao.DB.Table("file").
		Joins("JOIN file_tag_relation ON file.id = file_tag_relation.file_id").
		Where("file_tag_relation.tag_id = ? AND file_tag_relation.user_id = ?", tagID, userID).
		Find(&files).Error
	return files, err
}

// Delete 删除文件标签关联
func (dao *FileTagRelationDAO) Delete(fileID, tagID, userID uint64) error {
	return dao.DB.Where("file_id = ? AND tag_id = ? AND user_id = ?", fileID, tagID, userID).
		Delete(&models.FileTagRelation{}).Error
}

// DeleteByFileID 删除文件的所有标签关联
func (dao *FileTagRelationDAO) DeleteByFileID(fileID, userID uint64) error {
	return dao.DB.Where("file_id = ? AND user_id = ?", fileID, userID).
		Delete(&models.FileTagRelation{}).Error
}

// DeleteByTagID 删除标签的所有文件关联
func (dao *FileTagRelationDAO) DeleteByTagID(tagID, userID uint64) error {
	return dao.DB.Where("tag_id = ? AND user_id = ?", tagID, userID).
		Delete(&models.FileTagRelation{}).Error
}
