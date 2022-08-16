package user

import (
	"api_platforma/src/pkg/logging"
	"api_platforma/src/pkg/validation"
	"embed"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetUsers(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
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

func (h *service) CreateUser(c *fiber.Ctx) error {
	user := new(CreateUserDto)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	errors := validation.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	return c.JSON(user)
}
