package user

import (
	"api_platforma/src/pkg/logging"
	"embed"
	"github.com/kataras/iris/v12"
)

type Service interface {
	GetUsers(ctx iris.Context)
	CreateUser(ctx iris.Context)
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

func (h *service) GetUsers(c iris.Context) {
	fileData, err := h.embedFs.ReadFile("assets/fake.json")
	if err != nil {
		h.log.Warn(err)
	}

	c.JSON(fileData).Error()
}

func (h *service) CreateUser(c iris.Context) {
	user := new(CreateUserDto)

	user.Name = c.Params().Get("name")
	user.Password = c.Params().Get("password")

	c.JSON(user).Error()
}
