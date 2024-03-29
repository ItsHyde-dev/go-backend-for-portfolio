package projectsModel

import (
	"database/sql"

	"hyde.portfolio/database"
)

func Create(project *Project) (sql.Result, error) {
	query := `
    INSERT INTO projects (title, description, url)
    VALUES (
      :title, :description, :url
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
