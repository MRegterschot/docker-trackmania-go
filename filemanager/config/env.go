package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MRegterschot/docker-trackmania-plus/filemanager/structs"
	"github.com/joho/godotenv"
)

var AppEnv *structs.Env

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("FM_PORT"))
	if err != nil {
		port = 3300
	}

	userDataPath := os.Getenv("FM_USERDATA_PATH")
	if userDataPath == "" {
		userDataPath = "/server/UserData"
	}

	AppEnv = &structs.Env{
		Port:         port,
		LogLevel:     os.Getenv("FM_LOG_LEVEL"),
		UserDataPath: userDataPath,
	}

	return nil
}
