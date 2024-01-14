package projectsModel

import (
	"database/sql"

	"hyde.portfolio/database"
)

func create(project *Project) (sql.Result, error) {
	query := `
    INSERT INTO projects (title, description, url, technologies)
    VALUES (
      :title, :description, :url, :technologies
    )
  `

	return database.GetConnection().NamedExec(query, project)
}
