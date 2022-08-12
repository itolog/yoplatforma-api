package main

import (
	"api_platforma/src/internal/app"
	"api_platforma/src/internal/config"
)

func main() {
	cfg := config.NewConfig()
	server := app.NewApp(cfg)
	server.Logging.Info("Start server on port: ", cfg.Port)

	if err := server.Start(); err != nil {
		server.Logging.Fatal(err)
	}
}
