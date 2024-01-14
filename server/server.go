package server

import (
	"github.com/gofiber/fiber/v2"
	"hyde.portfolio/components/metadata"
)

func Init() {
	app := fiber.New();
  LoadRoutes(app)

  app.Listen(":3000")
}

func LoadRoutes(app *fiber.App) {
  metadata.Routes(app)
}
