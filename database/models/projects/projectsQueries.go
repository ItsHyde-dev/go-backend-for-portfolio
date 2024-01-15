package projectsModel

import (
	"database/sql"

	"hyde.portfolio/database"
)

func Create(project *Project) (sql.Result, error) {
	query := `
    INSERT INTO projects (title, description, url, technologies)
    VALUES (
      :title, :description, :url, :technologies
    )
  `

	return database.GetConnection().NamedExec(query, project)
}

func FindAll() ([]Project, error) {
  var projects []Project

  query := `
    SELECT * from projects
  `

  err := database.GetConnection().Select(&projects, query)

  return projects, err
}
