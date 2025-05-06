package main

import "github.com/MRegterschot/docker-trackmania-plus/filemanager/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
