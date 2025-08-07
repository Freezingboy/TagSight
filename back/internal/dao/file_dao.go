package dao

import "gorm.io/gorm"

type FileDao struct {
	db *gorm.DB
}

func NewFileDao(db *gorm.DB) *FileDao {
	return &FileDao{db: db}
}
