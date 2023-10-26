package config

import (
	// "database/sql"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {

	dsn := os.Getenv("DB_PG_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Failed, Err : ", err.Error())
	}

	return db

	// MySQL
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// dbString := os.Getenv("DB_STRING")

	// db, err := sql.Open("mysql", dbString)
	// if err != nil {
	// 	log.Fatal("DB Conection Failed")
	// }
}
