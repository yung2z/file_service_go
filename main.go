package main

import (
	"files_service/controllers"
	"files_service/repositories"
	"files_service/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	filesRepository := repositories.NewFilesRepository()
	filesService := services.NewFilesService(filesRepository)

	controllers.NewFilesController(app, filesService)

	app.Static("/static", "./static")

	app.Listen(":8080")
}