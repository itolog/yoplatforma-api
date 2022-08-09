package app

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/user"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	server *fiber.App
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
		server: fiber.New(fiber.Config{
			Prefork:     true,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (app *App) Start() error {
	userHandler := user.NewUserHandler()
	userHandler.Register(app.server)

	return app.server.Listen(app.config.Port)
}
