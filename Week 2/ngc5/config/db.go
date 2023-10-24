package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConn() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbString := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", dbString)
	if err != nil {
		log.Fatal("DB connection failed")
	}

	return db
}
