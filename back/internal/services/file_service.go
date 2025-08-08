package services

import (
	"back/api/response"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"back/internal/dao"
	"back/internal/models"
)

var (
	// ErrFileNotFound 文件不存在错误
	ErrFileNotFound = errors.New("文件不存在")
	// ErrFilePathExists 文件路径已存在错误
	ErrFilePathExists = errors.New("文件路径已存在")
)

// FileService 文件服务
type FileService struct {
	fileDAO            *dao.FileDAO
	fileTagRelationDAO *dao.FileTagRelationDAO
	tagDAO             *dao.TagDAO
	uploadDir          string
}

// NewFileService 创建FileService实例
func NewFileService(fileDAO *dao.FileDAO, fileTagRelationDAO *dao.FileTagRelationDAO, tagDAO *dao.TagDAO, uploadDir string) *FileService {
	return &FileService{
		fileDAO:            fileDAO,
		fileTagRelationDAO: fileTagRelationDAO,
		tagDAO:             tagDAO,
		uploadDir:          uploadDir,
	}
}

// UploadFile 上传多个文件
func (s *FileService) UploadFile(files []*multipart.FileHeader, userID uint64) ([]*response.FileResponse, error) {
	var responses []*response.FileResponse

	// 创建用户上传目录
	userUploadDir := filepath.Join(s.uploadDir, fmt.Sprintf("user%d", userID))
	if err := os.MkdirAll(userUploadDir, 0755); err != nil {
		return nil, err
	}

	// 遍历所有上传的文件
	for _, fileHeader := range files {
		// 打开文件
		src, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open uploaded file: %w", err)
		}
		defer src.Close()

		// 获取文件名和扩展
		fileName := fileHeader.Filename
		fileExt := filepath.Ext(fileName)
		fileType := strings.TrimPrefix(fileExt, ".")
		if fileType == "" {
			fileType = "unknown"
		}

		// 使用时间戳确保文件名唯一
		timestamp := time.Now().UnixNano()
		storedFileName := fmt.Sprintf("%d_%s", timestamp, fileName)
		filePath := filepath.Join(userUploadDir, storedFileName)
		relativePath := filepath.Join("/static", fmt.Sprintf("user%d", userID), storedFileName)

		// 检查文件路径是否已存在
		existingFile, err := s.fileDAO.GetByPath(relativePath)
		if err != nil {
			return nil, err
		}
		if existingFile != nil {
			return nil, fmt.Errorf("file already exists: %s", relativePath)
		}

		// 创建目标文件
		dst, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}

		// 拷贝内容
		if _, err = io.Copy(dst, src); err != nil {
			dst.Close()
			return nil, err
		}
		dst.Close()

		// 创建数据库记录
		file := &models.File{
			UserID:     userID,
			Name:       fileName,
			Size:       int(fileHeader.Size),
			Type:       fileType,
			Path:       relativePath,
			UploadTime: time.Now(),
		}
		if err := s.fileDAO.Create(file); err != nil {
			// 删除已上传的文件
			_ = os.Remove(filePath)
			return nil, err
		}

		// 构建响应
		resp := &response.FileResponse{
			ID:         file.ID,
			Name:       file.Name,
			Size:       file.Size,
			Type:       file.Type,
			Path:       file.Path,
			UploadTime: file.UploadTime,
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

// GetFileList 获取文件列表
func (s *FileService) GetFileList(userID uint64, page, size int, tagID uint64) (*response.FileListResponse, error) {
	var files []models.File
	var total int64
	var err error

	// 根据标签ID筛选文件
	if tagID > 0 {
		files, total, err = s.fileDAO.ListByTagID(userID, tagID, page, size)
	} else {
		files, total, err = s.fileDAO.ListByUserID(userID, page, size)
	}
	if err != nil {
		return nil, err
	}

	// 构建响应
	res := &response.FileListResponse{
		Total: total,
		List:  make([]response.FileResponse, 0, len(files)),
	}

	// 填充文件信息
	for _, file := range files {
		fileResp := response.FileResponse{
			ID:         file.ID,
			Name:       file.Name,
			Size:       file.Size,
			Type:       file.Type,
			UploadTime: file.UploadTime,
			Tags:       make([]response.TagResponse, 0),
		}

		// 查询文件的标签
		tags, err := s.fileTagRelationDAO.ListTagsByFileID(file.ID, userID)
		if err != nil {
			return nil, err
		}

		// 填充标签信息
		for _, tag := range tags {
			fileResp.Tags = append(fileResp.Tags, response.TagResponse{
				ID:    tag.ID,
				Name:  tag.Name,
				Color: tag.Color,
			})
		}

		res.List = append(res.List, fileResp)
	}

	return res, nil
}

// GetFile 获取文件
func (s *FileService) GetFile(fileID, userID uint64) (*models.File, error) {
	file, err := s.fileDAO.GetByID(fileID, userID)
	if err != nil {
		return nil, err
	}
	if file == nil {
		return nil, ErrFileNotFound
	}
	return file, nil
}

// DeleteFile 删除文件
func (s *FileService) DeleteFile(fileID, userID uint64) error {
	// 查询文件
	file, err := s.fileDAO.GetByID(fileID, userID)
	if err != nil {
		return err
	}
	if file == nil {
		return ErrFileNotFound
	}

	// 删除文件标签关联
	if err := s.fileTagRelationDAO.DeleteByFileID(fileID, userID); err != nil {
		return err
	}

	// 删除文件记录
	if err := s.fileDAO.Delete(fileID, userID); err != nil {
		return err
	}

	// 删除物理文件
	filePath := filepath.Join(".", file.Path)

	return os.Remove(filePath)
}
