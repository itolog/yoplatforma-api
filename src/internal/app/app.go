package app

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/user"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"os"
)

type App struct {
	server *fiber.App
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
		server: fiber.New(fiber.Config{
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

	port := os.Getenv("PORT")

	return app.server.Listen(":" + port)
}
