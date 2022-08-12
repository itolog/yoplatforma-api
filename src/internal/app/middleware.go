package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
	"path/filepath"
)

func configureApp(app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())

	p, err := filepath.Abs("src/assets/fake.json")

	if err != nil {

		log.Fatal(err)
	}
	fmt.Println(p)
	fmt.Printf("Base: %s\n", filepath.Base(p))
	fmt.Printf("Dir: %s\n", filepath.Dir(p))

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir(p),
	}))
}
