package handler

import (
	"avengers-v2/config"
	"avengers2/entity"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GetVillain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := config.DBConn()
	defer db.Close()

	w.Header().Set("content-type", "application/json")

	// Villain Object
	villain := []entity.Villain{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
	SELECT name, universe, imageURL FROM villain
	`)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Failed when retrive data villain.",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var v entity.Villain

		err := rows.Scan(&v.Name, &v.Universe, &v.ImageURL)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"msg": "Failed when retrive data villain.",
			})
		}

		villain = append(villain, v)
	}

	json.NewEncoder(w).Encode(villain)
}
