package services

import (
	"back/api/request"
	"back/api/response"
	"errors"

	"back/internal/dao"
	"back/internal/models"
)

var (
	// ErrFileTagRelationExists 文件标签关联已存在错误
	ErrFileTagRelationExists = errors.New("文件已添加该标签")
	// ErrFileTagRelationNotFound 文件标签关联不存在错误
	ErrFileTagRelationNotFound = errors.New("文件标签关联不存在")
)

// FileTagService 文件标签关联服务
type FileTagService struct {
	fileTagRelationDAO *dao.FileTagRelationDAO
	fileDAO            *dao.FileDAO
	tagDAO             *dao.TagDAO
}

// NewFileTagService 创建FileTagService实例
func NewFileTagService(fileTagRelationDAO *dao.FileTagRelationDAO, fileDAO *dao.FileDAO, tagDAO *dao.TagDAO) *FileTagService {
	return &FileTagService{
		fileTagRelationDAO: fileTagRelationDAO,
		fileDAO:            fileDAO,
		tagDAO:             tagDAO,
	}
}

// AddTagToFile 为文件添加标签
func (s *FileTagService) AddTagToFile(req *request.FileTagRequest, userID uint64) (*response.FileTagResponse, error) {
	// 检查文件是否存在且属于当前用户
	file, err := s.fileDAO.GetByID(req.FileID, userID)
	if err != nil {
		return nil, err
	}
	if file == nil {
		return nil, ErrFileNotFound
	}

	// 检查标签是否存在且属于当前用户
	tag, err := s.tagDAO.GetByID(req.TagID, userID)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, ErrTagNotFound
	}

	// 检查关联是否已存在
	relation, err := s.fileTagRelationDAO.GetByFileAndTag(req.FileID, req.TagID, userID)
	if err != nil {
		return nil, err
	}
	if relation != nil {
		return nil, ErrFileTagRelationExists
	}

	// 创建关联
	newRelation := &models.FileTagRelation{
		UserID: userID,
		FileID: req.FileID,
		TagID:  req.TagID,
	}
	if err := s.fileTagRelationDAO.Create(newRelation); err != nil {
		return nil, err
	}

	// 返回响应
	return &response.FileTagResponse{
		ID:     newRelation.ID,
		FileID: newRelation.FileID,
		TagID:  newRelation.TagID,
	}, nil
}

// RemoveTagFromFile 从文件移除标签
func (s *FileTagService) RemoveTagFromFile(req *request.FileTagRequest, userID uint64) error {
	// 检查关联是否存在
	relation, err := s.fileTagRelationDAO.GetByFileAndTag(req.FileID, req.TagID, userID)
	if err != nil {
		return err
	}
	if relation == nil {
		return ErrFileTagRelationNotFound
	}

	// 删除关联
	return s.fileTagRelationDAO.Delete(req.FileID, req.TagID, userID)
}

// GetFileTags 获取文件的标签列表
func (s *FileTagService) GetFileTags(fileID, userID uint64) []response.TagResponse {
	// 检查文件是否存在且属于当前用户
	file, err := s.fileDAO.GetByID(fileID, userID)
	if err != nil {
		return nil
	}
	if file == nil {
		return nil
	}

	// 查询文件的标签
	tags, err := s.fileTagRelationDAO.ListTagsByFileID(fileID, userID)
	if err != nil {
		return nil
	}

	// 转换为响应格式
	var res []response.TagResponse
	for _, tag := range tags {
		res = append(res, response.TagResponse{
			ID:    tag.ID,
			Name:  tag.Name,
			Color: tag.Color,
		})
	}

	return res
}
