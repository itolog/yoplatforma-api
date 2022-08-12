package user

import (
	"api_platforma/src/internal/handlers"
	"api_platforma/src/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"log"
	"path/filepath"
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
	p, err := filepath.Abs("src/assets/fake.json")

	if err != nil {

		log.Fatal(err)
	}
	c.Status(fiber.StatusOK)
	return c.SendFile(p, true)
}
