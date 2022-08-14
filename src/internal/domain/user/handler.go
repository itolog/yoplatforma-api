package user

import (
	"api_platforma/src/internal/handlers"
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/gofiber/fiber/v2"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
)

type handler struct {
	log     *logging.Logger
	embedFs *embed.FS
}

func NewUserHandler(log *logging.Logger, embedFs *embed.FS) handlers.Handler {
	return &handler{
		log:     log,
		embedFs: embedFs,
	}
}

func (h *handler) Register(app *fiber.App) {
	app.Get(usersUrl, h.GetUsers)
}

func (h *handler) GetUsers(c *fiber.Ctx) error {

	fileData, err := h.embedFs.ReadFile("assets/fake.json")
	if err != nil {
		h.log.Warn(err)
	}

	c.Status(fiber.StatusOK)
	c.Type("json")
	return c.Send(fileData)
}
