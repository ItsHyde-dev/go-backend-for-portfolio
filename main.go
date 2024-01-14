package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"hyde.portfolio/database"
	"hyde.portfolio/server"
)

func main () {

  godotenv.Load()

  db := database.GetConnection()
  defer db.Close()
  defer fmt.Println("Database connection closed!")

  server.Init()
}
