package user

import (
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetUsers(ctx *fiber.Ctx) error
}

type service struct {
	log     *logging.Logger
	embedFs *embed.FS
	storage Storage
}

func NewService(log *logging.Logger, embedFs *embed.FS) Service {
	return &service{
		log:     log,
		embedFs: embedFs,
	}
}

func (h *service) GetUsers(c *fiber.Ctx) error {
	fileData, err := h.embedFs.ReadFile("assets/fake.json")
	if err != nil {
		h.log.Warn(err)
	}

	c.Status(fiber.StatusOK)
	c.Type("json")
	return c.Send(fileData)
}
