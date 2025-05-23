package app

import (
	"github.com/MRegterschot/docker-trackmania-plus/filemanager/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/UserData/*", handlers.HandleListFiles)
	app.Post("/upload", handlers.HandleUploadFiles)
	app.Delete("/delete", handlers.HandleDeleteFiles)
}
