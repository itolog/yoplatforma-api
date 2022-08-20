package user

import (
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/kataras/iris/v12/core/router"
)

const (
	usersUrl = "/users/"
)

func NewUserController(log *logging.Logger, embedFs *embed.FS, app router.Party) {
	userService := NewService(log, embedFs)
	// [/api/v1/users] Get all users
	app.Get(usersUrl, userService.GetUsers)

	app.Post(usersUrl, userService.CreateUser)
}
