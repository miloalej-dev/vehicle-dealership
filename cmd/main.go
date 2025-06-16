package main

import (
	"fmt"
	"github.com/miloalej-dev/vehicle-dealership/cmd/server"
	"os"
)

func main() {
	// env
	// ...
	serverAddress := os.Getenv("SERVER_ADDRESS")
	loaderFile := os.Getenv("LOADER_FILE")

	// app
	// - config
	cfg := &server.ConfigServerChi{
		ServerAddress:  serverAddress,
		LoaderFilePath: loaderFile,
	}
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
