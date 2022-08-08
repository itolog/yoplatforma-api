package app

import (
	"api_platforma/src/internal/user"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	server *fiber.App
	config *Config
}

func newApp(config *Config) *App {
	return &App{
		config: config,
		server: fiber.New(fiber.Config{
			Prefork:     true,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func Start() error {
	config := newConfig()
	app := newApp(config)

	userHandler := user.NewUserHandler()
	userHandler.Register(app.server)

	return app.server.Listen(config.Port)
}
