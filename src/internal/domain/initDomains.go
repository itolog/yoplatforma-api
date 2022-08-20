package domain

import (
	"api_platforma/src/internal/config"
	"api_platforma/src/internal/domain/user"
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/kataras/iris/v12"
)

type Option struct {
	Server  *iris.Application
	Logging *logging.Logger
	EmbedFs *embed.FS
	Config  *config.Config
}

func InitDomains(app Option) {
	v1 := app.Server.Party(app.Config.PrefixV1)
	v1.Get("sa", func(context iris.Context) {

	})
	user.NewUserController(app.Logging, app.EmbedFs, v1)
}
