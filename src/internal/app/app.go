package app

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/domain"
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type App struct {
	server  *fiber.App
	config  *config.Config
	Logging *logging.Logger
	embedFs *embed.FS
}

func NewApp(config *config.Config, embedFs *embed.FS) *App {
	return &App{
		config: config,
		server: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		Logging: logging.GetLogger(),
		embedFs: embedFs,
	}
}

func (app *App) Start() error {
	// Init Middleware
	configureApp(app.server, app.embedFs, app.config)

	app.server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
	app.server.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	domain.InitDomains(domain.Option{
		Server:  app.server,
		Logging: app.Logging,
		EmbedFs: app.embedFs,
		Config:  app.config,
	})

	return app.server.Listen(":" + app.config.Port)
}
