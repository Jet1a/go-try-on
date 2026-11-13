package server

import (
	"github.com/Jet1a/go-try-on/api/internal/config"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.Config) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
