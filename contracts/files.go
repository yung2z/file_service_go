package contracts

import "mime/multipart"

type Res struct {
	Message string `json:"message"`
}

type CreateFileRequest struct {
	Name     string
	File multipart.FileHeader
}

type CreateFileResponse struct{
	Id int	`json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}