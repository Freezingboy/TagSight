package services

import (
	"back/api/request"
	"back/api/response"
	"errors"
	"gorm.io/gorm"

	"back/internal/dao"
	"back/internal/models"
)

var (
	// ErrTagExists 标签已存在错误
	ErrTagExists = errors.New("标签名已存在")
	// ErrTagNotFound 标签不存在错误
	ErrTagNotFound = errors.New("标签不存在")
)

// TagService 标签服务
type TagService struct {
	tagDAO             *dao.TagDAO
	fileTagRelationDAO *dao.FileTagRelationDAO
}

// NewTagService 创建TagService实例
func NewTagService(tagDAO *dao.TagDAO, fileTagRelationDAO *dao.FileTagRelationDAO) *TagService {
	return &TagService{
		tagDAO:             tagDAO,
		fileTagRelationDAO: fileTagRelationDAO,
	}
}

// CreateTag 创建标签
func (s *TagService) CreateTag(req *request.CreateTagRequest, userID uint64) (*response.CreateTagResponse, error) {
	// 检查标签名是否已存在
	existingTag, err := s.tagDAO.GetByName(req.Name, userID)
	if err != nil {
		return nil, err
	}
	if existingTag != nil {
		return nil, ErrTagExists
	}

	// 创建标签
	tag := &models.Tag{
		UserID: userID,
		Name:   req.Name,
		Color:  req.Color,
	}
	if err := s.tagDAO.Create(tag); err != nil {
		return nil, err
	}

	// 返回响应
	return &response.CreateTagResponse{
		ID:     tag.ID,
		UserID: tag.UserID,
		Name:   tag.Name,
		Color:  tag.Color,
	}, nil
}

// GetUserTags 获取用户的所有标签
func (s *TagService) GetUserTags(userID uint64) ([]response.TagResponse, error) {
	// 查询标签
	tags, err := s.tagDAO.ListByUserID(userID)
	if err != nil {
		return nil, err
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

	return res, nil
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	Name  string `json:"name" binding:"required,max=50"`
	Color string `json:"color" binding:"required,max=20"`
}

// UpdateTag 更新标签
func (s *TagService) UpdateTag(tagID uint64, req *UpdateTagRequest, userID uint64) (*response.TagResponse, error) {
	// 查询标签
	tag, err := s.tagDAO.GetByID(tagID, userID)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, ErrTagNotFound
	}

	// 检查新名称是否与其他标签冲突
	if tag.Name != req.Name {
		existingTag, err := s.tagDAO.GetByName(req.Name, userID)
		if err != nil {
			return nil, err
		}
		if existingTag != nil && existingTag.ID != tagID {
			return nil, ErrTagExists
		}
	}

	// 更新标签
	tag.Name = req.Name
	tag.Color = req.Color
	if err := s.tagDAO.Update(tag); err != nil {
		return nil, err
	}

	// 返回响应
	return &response.TagResponse{
		ID:    tag.ID,
		Name:  tag.Name,
		Color: tag.Color,
	}, nil
}

// DeleteTag 删除标签
func (s *TagService) DeleteTag(tagID uint64, userID uint64) error {
	// 查询标签
	tag, err := s.tagDAO.GetByID(tagID, userID)
	if err != nil {
		return err
	}
	if tag == nil {
		return ErrTagNotFound
	}

	// 开启事务
	err = s.tagDAO.DB.Transaction(func(tx *gorm.DB) error {
		// 删除标签关联
		if err := s.fileTagRelationDAO.DeleteByTagID(tagID, userID); err != nil {
			return err
		}

		// 删除标签
		if err := s.tagDAO.Delete(tagID, userID); err != nil {
			return err
		}

		return nil
	})

	return err
}
