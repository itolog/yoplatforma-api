package user

import (
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/gofiber/fiber/v2"
)

const (
	usersUrl = "/users/"
)

func NewUserController(log *logging.Logger, embedFs *embed.FS, app fiber.Router) {
	userService := NewService(log, embedFs)
	// [/api/v1/users] Get all users
	app.Get(usersUrl, userService.GetUsers)
}
