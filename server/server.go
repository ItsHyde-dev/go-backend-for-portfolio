package server

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"hyde.portfolio/components/metadata"
)

func Init() {
	app := fiber.New();
  LoadRoutes(app)

  port := os.Getenv("PORT")
  if strings.TrimSpace(port) == "" {
    port = ":3000"
  }

  app.Listen(port)
}

func LoadRoutes(app *fiber.App) {
  metadata.Routes(app)
}
