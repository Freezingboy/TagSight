package services

import (
	"back/internal/dao"
)

type IFileService interface {
}

type FileServiceImpl struct {
	fileDao *dao.FileDao
}

func NewFileService(fileDao *dao.FileDao) IFileService {
	return &FileServiceImpl{fileDao: fileDao}
}
