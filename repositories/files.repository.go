package repositories

import (
	"files_service/contracts"
	"files_service/models"
)

type FilesRepository struct {
	files []models.FileModel
}

func NewFilesRepository() *FilesRepository {
	return &FilesRepository{files: []models.FileModel{}}
}

func (r *FilesRepository) GetAll() []models.FileModel{
	return r.files
}

func (r *FilesRepository) Create(data contracts.CreateFileRequest, link string) models.FileModel {
	newFile := models.FileModel{Id: len(r.files) + 1, Name:  data.Name, Link: link}
	r.files = append(r.files, newFile)
	return newFile
}