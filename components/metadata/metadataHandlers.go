package metadata

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"hyde.portfolio/utils"
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

  err := utils.ValidateStruct(body);
  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "message": err.Error(),
    })
  }

  fmt.Println(body)

  return c.JSON(fiber.Map{
    "message": "Added new project",
  })
}
