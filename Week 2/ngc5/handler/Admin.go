package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Admin struct {
	DB *sql.DB
}

func NewAdminHandler(db *sql.DB) Admin {
	return Admin{
		DB: db,
	}
}

func (handler Admin) GetAdmin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"msg": "Hello Admin",
	})
}
