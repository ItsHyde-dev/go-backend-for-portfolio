package metadata

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) *fiber.Router {
  router := app.Group("/metadata")

  router.Get("/", getMetadata)
  router.Post("/addProject", addProject)

  return &router;
}
