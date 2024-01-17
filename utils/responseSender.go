package utils

import "github.com/gofiber/fiber/v2"

func SendResponse(
  c *fiber.Ctx, status int, message string, data interface{},
) error {

	response := fiber.Map{
		"message": message,
	}

	if data != nil {
		response["data"] = data
	}

	return c.Status(status).JSON(response)
}
