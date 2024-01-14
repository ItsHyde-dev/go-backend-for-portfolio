package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func GetConnection() *sqlx.DB {

	fmt.Println("Getting database connection...")

	if Db != nil {
		fmt.Println("Using existing database connection...")
		return Db
	}

	fmt.Println("Returning new connection...")
	return Connect()
}

func Connect() *sqlx.DB {

	fmt.Println("Connecting to database...")

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := sqlx.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	Db = db
	return db
}
