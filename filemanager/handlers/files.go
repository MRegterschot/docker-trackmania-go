package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/MRegterschot/docker-trackmania-plus/filemanager/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Handle files upload
func HandleUploadFiles(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid form data")
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No files found")
	}

	var errors []string

	for _, file := range files {
		dir := filepath.Dir(file.Filename)

		// Set the destination path for the uploaded file
		dest := filepath.Join(config.AppEnv.UserDataPath, filepath.Clean("/"+dir))

		// Check if the path is in the UserData directory
		if !strings.HasPrefix(dest, config.AppEnv.UserDataPath) {
			errors = append(errors, "Invalid file path: "+file.Filename)
			continue
		}

		// Create the directory if it doesn't exist
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			zap.L().Error("Error creating directory", zap.String("path", dest), zap.Error(err))
			errors = append(errors, "Failed to create directory: "+dest)
			continue
		}

		// Save the file to the destination
		targetPath := filepath.Join(dest, filepath.Base(file.Filename))
		if err := c.SaveFile(file, targetPath); err != nil {
			zap.L().Error("Error saving file", zap.String("path", targetPath), zap.Error(err))
			errors = append(errors, "Failed to save file: "+file.Filename)
			continue
		}

		zap.L().Info("File uploaded", zap.String("path", targetPath))
	}

	if len(errors) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Some files could not be uploaded",
			"errors":  errors,
		})
	}

	return c.SendString("Files uploaded successfully")
}

// Handle file deletion
func HandleDeleteFiles(c *fiber.Ctx) error {
	// Get the file path from the request
	var paths []string
	if err := c.BodyParser(&paths); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if len(paths) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No file paths provided")
	}

	var errors []string

	for _, path := range paths {
		// Set the destination path for the uploaded file
		cleanPath := filepath.Join(config.AppEnv.UserDataPath, filepath.Clean("/"+path))

		// Check if the path is in the UserData directory
		if !strings.HasPrefix(cleanPath, config.AppEnv.UserDataPath) {
			errors = append(errors, "Invalid file path: "+path)
			continue
		}

		// Check if the file exists before trying to delete it
		if _, err := os.Stat(cleanPath); err != nil {
			if os.IsNotExist(err) {
				errors = append(errors, "File does not exist: "+path)
				continue
			}
			zap.L().Error("Error checking file existence", zap.Error(err))
			errors = append(errors, "Error checking file existence: "+path)
			continue
		}

		// Delete the file
		if err := os.Remove(cleanPath); err != nil {
			zap.L().Error("Error deleting file", zap.String("path", cleanPath), zap.Error(err))
			errors = append(errors, "Failed to delete file: "+path)
			continue
		}

		zap.L().Info("File deleted", zap.String("path", cleanPath))
	}

	if len(errors) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Some files could not be deleted",
			"errors":  errors,
		})
	}

	return c.SendString("Files deleted successfully")
}
