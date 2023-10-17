package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (*sql.DB, error) {
	dbUrl := "root:@tcp(127.0.0.1:3306)/avengers"

	db, err := sql.Open("mysql", dbUrl)
	return db, err
}
