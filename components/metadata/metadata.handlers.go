package metadata

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"hyde.portfolio/database/models/projects"
	"hyde.portfolio/utils"
)

func getMetadata(c *fiber.Ctx) error {
	fmt.Println("Hello, World!")

	return utils.SendResponse(c, 200, "Hello, World!", nil)
}

func addProject(c *fiber.Ctx) error {

	var body AddProject
	c.BodyParser(&body)

	err := utils.ValidateStruct(body)
	if err != nil {
		return utils.SendResponse(c, 400, err.Error(), nil)
	}

	result, err := projectsModel.Create(&projectsModel.Project{
		Title:        body.Name,
		Url:          body.Url,
		Description:  body.Description,
		Technologies: body.Technologies,
	})

	if err != nil {
		fmt.Println(err)
		return utils.SendResponse(
			c, 400, "Could not create the project", nil,
		)
	}

	fmt.Println(result.LastInsertId())

  return utils.SendResponse(c, 200, "Project created", nil)
}

func getProjects(c *fiber.Ctx) error {

	fmt.Println("Getting Projects")

	projects, err := projectsModel.FindAll()
	if err != nil {
    return utils.SendResponse(c, 400, err.Error(), nil)
	}

  return utils.SendResponse(c, 200, "Successfully Fetched Projects", projects)
}
