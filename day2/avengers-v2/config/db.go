package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() *sql.DB {
	dbUrl := "root:@tcp(127.0.0.1:3306)/avengers"

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to DB")
	}

	return db
}
