package user

import (
	"api_platforma/src/internal/handlers"
	"api_platforma/src/pkg/logging"
	"github.com/gofiber/fiber/v2"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
)

type handler struct {
	log *logging.Logger
}

func NewUserHandler(log *logging.Logger) handlers.Handler {
	return &handler{
		log: log,
	}
}

func (h *handler) Register(app *fiber.App) {
	app.Get(usersUrl, h.GetUsers)
}

func (h *handler) GetUsers(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendFile("src/assets/fake.json", true)
}
