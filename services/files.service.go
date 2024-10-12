package services

import (
	"files_service/contracts"
	"files_service/models"
	"files_service/repositories"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type FilesServices struct {
	filesRepository *repositories.FilesRepository
}

func NewFilesService(filesRepository *repositories.FilesRepository) *FilesServices{
	return &FilesServices{filesRepository: filesRepository}
}

func (s *FilesServices) UploadFile(ctx *fiber.Ctx, data contracts.CreateFileRequest) (models.FileModel, error){
	
	randStr := uuid.New().String()
	mimetype := strings.Split(data.File.Filename, ".")[1]

	fileName := randStr+"."+mimetype
	dirName := "static"

	dirPath := filepath.Join(".", dirName)

	if _, existErr := os.Stat(dirPath); os.IsNotExist(existErr) {
		os.Mkdir(dirPath, os.ModePerm)
	}
	
	filePath := filepath.Join(dirPath, fileName)
	creationErr := ctx.SaveFile(&data.File, filePath)

	if creationErr != nil {
		return models.FileModel{}, creationErr
	}

	link := dirName + "/" + fileName

	createdFile := s.filesRepository.Create(data, link)
	return createdFile, nil
}

func (s *FilesServices) GetAll() []models.FileModel{
	return s.filesRepository.GetAll()
}
