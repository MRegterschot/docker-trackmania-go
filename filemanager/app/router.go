package app

import (
	"github.com/MRegterschot/docker-trackmania-plus/filemanager/config"
	"github.com/MRegterschot/docker-trackmania-plus/filemanager/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/upload", handlers.HandleUploadFiles)
	app.Delete("/delete", handlers.HandleDeleteFiles)

	// Setup static file serving
	app.Static("/UserData", config.AppEnv.UserDataPath)
}
