package metadata

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getMetadata(c *fiber.Ctx) error {
  fmt.Println("Hello, World!")

  return c.JSON(fiber.Map{
    "message": "Hello, World!",
  })
}

func addProject(c *fiber.Ctx) error {

  var body AddProject;
  c.BodyParser(&body)

  fmt.Println(body)

  return c.JSON(fiber.Map{
    "message": "Added new project",
  })
}
