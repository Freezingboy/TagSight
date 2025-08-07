package controllers

import (
	"back/internal/services"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	service services.IFileService
}

func NewFileController(service services.IFileService) *FileController {
	return &FileController{
		service: service,
	}
}

func (fc *FileController) UploadFile(ctx *gin.Context) {

}
