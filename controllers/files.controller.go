package controllers

import (
	"files_service/contracts"
	"files_service/services"

	"github.com/gofiber/fiber/v2"
)

type FilesController struct {
	filesService *services.FilesServices
}

func NewFilesController(app *fiber.App, filesService *services.FilesServices) *FilesController {
	controller := &FilesController{filesService: filesService}
	app.Post("/files", controller.CreateFile)
	app.Get("/files", controller.GetFiles)
	return controller
}

func (c *FilesController) GetFiles(ctx *fiber.Ctx) error{
	response := c.filesService.GetAll()
	return ctx.JSON(response)
}

func (c *FilesController) CreateFile(ctx *fiber.Ctx) error{
	file, _ := ctx.FormFile("file")
	name := ctx.FormValue("name")

	req := contracts.CreateFileRequest{Name: name, File: *file}

	createdFile, _ := c.filesService.UploadFile(ctx, req)
		
	response := contracts.CreateFileResponse{Id: createdFile.Id, Name: createdFile.Name, Link: createdFile.Link}

	return ctx.JSON(response)
}