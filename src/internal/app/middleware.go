package app

import (
	"api_platforma/src/internal/config"
	"embed"
	"github.com/kataras/iris/v12"
)

func configureApp(app *iris.Application, embedFs *embed.FS, config *config.Config) {
	//app.Use(recover.New())
	//app.Use(logger.New())
	//app.Use(cors.New())
	//app.Use(compress.New())
	//
	//app.Use(config.PrefixV1, filesystem.New(filesystem.Config{
	//	Root: http.FS(embedFs),
	//}))
}
