package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// FileUtil 文件工具类
type FileUtil struct{}

// SaveUploadedFile 保存上传的文件到指定目录
func (fu *FileUtil) SaveUploadedFile(file *multipart.FileHeader, userId uint64) (string, error) {
	// 创建static目录（如果不存在）
	staticDir := "static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		err = os.Mkdir(staticDir, 0755)
		if err != nil {
			return "", fmt.Errorf("创建static目录失败: %w", err)
		}
	}

	// 创建用户目录
	userDir := filepath.Join(staticDir, fmt.Sprintf("%d", userId))
	if _, err := os.Stat(userDir); os.IsNotExist(err) {
		err = os.Mkdir(userDir, 0755)
		if err != nil {
			return "", fmt.Errorf("创建用户目录失败: %w", err)
		}
	}

	// 构建文件路径
	dst := filepath.Join(userDir, file.Filename)

	// 检查文件是否已存在
	if _, err := os.Stat(dst); err == nil {
		return "", errors.New("文件已存在")
	}

	// 保存文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 读取文件内容并写入
	buffer := make([]byte, 1024*1024) // 1MB buffer
	for {
		n, err := src.Read(buffer)
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		if _, err := out.Write(buffer[:n]); err != nil {
			return "", err
		}
	}

	// 返回相对路径
	return "/" + dst, nil
}

// GetFileType 获取文件类型
func (fu *FileUtil) GetFileType(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return "unknown"
	}
	return strings.TrimPrefix(ext, ".")
}