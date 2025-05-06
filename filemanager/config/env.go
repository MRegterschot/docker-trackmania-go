package config

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MRegterschot/docker-trackmania-plus/filemanager/structs"
	"github.com/joho/godotenv"
)

var AppEnv *structs.Env

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.New("failed to load .env file")
	}

	port, err := strconv.Atoi(os.Getenv("FM_PORT"))
	if err != nil {
		port = 3300
	}

	userDataPath := os.Getenv("FM_USERDATA_PATH")
	if userDataPath == "" {
		userDataPath = "/server/UserData"
	}

	absPath, err := filepath.Abs(userDataPath)
	if err != nil {
		return errors.New("failed to get absolute path for UserData directory")
	}

	AppEnv = &structs.Env{
		Port:         port,
		LogLevel:     os.Getenv("FM_LOG_LEVEL"),
		UserDataPath: absPath,
	}

	return nil
}
