package repository

import "database/sql"

type User struct {
	DB *sql.DB
}
