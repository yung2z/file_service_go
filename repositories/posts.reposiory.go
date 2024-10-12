package repositories

import (
	"files_service/models"
)

type PostsRepository struct {
	posts []models.PostsModel
}

func (r *PostsRepository) GetAll() []models.PostsModel{
	return r.posts
}

func (r *PostsRepository) Create() {
	
	
}