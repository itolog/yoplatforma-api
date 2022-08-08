package user

import (
	"api_platforma/src/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
)

type handler struct {
}

func NewUserHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(app *fiber.App) {
	app.Get(usersUrl, h.GetUsers)
}

func (h *handler) GetUsers(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendString("Users")
}
