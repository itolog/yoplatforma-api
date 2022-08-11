package app

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/user"
	"api_platforma/src/pkg/logging"
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

	app.server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	logging.Info(app.config.Port)
	return app.server.Listen(":8000")
}
