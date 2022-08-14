package domain

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/domain/user"
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/gofiber/fiber/v2"
)

type Option struct {
	Server  *fiber.App
	Logging *logging.Logger
	EmbedFs *embed.FS
	Config  *config.Config
}

func InitDomains(app Option) {
	v1 := app.Server.Group(app.Config.PrefixV1)

	user.NewUserController(app.Logging, app.EmbedFs, v1)
}
