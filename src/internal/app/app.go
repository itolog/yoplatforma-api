package app

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/domain"
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/kataras/iris/v12"
)

type App struct {
	server  *iris.Application
	config  *config.Config
	Logging *logging.Logger
	embedFs *embed.FS
}

func NewApp(config *config.Config, embedFs *embed.FS) *App {
	return &App{
		config: config,
		//server: fiber.New(fiber.Config{
		//	JSONEncoder: json.Marshal,
		//	JSONDecoder: json.Unmarshal,
		//}),
		server:  iris.New(),
		Logging: logging.GetLogger(),
		embedFs: embedFs,
	}
}

func (app *App) Start() error {
	// Init Middleware
	configureApp(app.server, app.embedFs, app.config)

	//app.server.Get("/", func(ctx iris.Context) error {
	//	return ctx.SendString("Hello World!")
	//})
	//app.server.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	domain.InitDomains(domain.Option{
		Server:  app.server,
		Logging: app.Logging,
		EmbedFs: app.embedFs,
		Config:  app.config,
	})

	return app.server.Listen(":" + app.config.Port)
}
