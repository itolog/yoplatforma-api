package app

import (
	"api_platforma/src/internal/config"
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
)

func configureApp(app *fiber.App, embedFs *embed.FS, config *config.Config) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())

	app.Use(config.PrefixV1, filesystem.New(filesystem.Config{
		Root: http.FS(embedFs),
	}))
}
