package main

import (
	"api_platforma/src/internal/app"
	"api_platforma/src/internal/config"
	"api_platforma/src/pkg/logging"
)

func main() {
	cfg := config.NewConfig()
	server := app.NewApp(cfg)
	logging.Info("Start server on port ", cfg.Port)
	if err := server.Start(); err != nil {
		logging.Fatal(err)
	}
}
