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

func GetHeroes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := config.DBConn()
	defer db.Close()

	w.Header().Set("content-type", "application/json")

	// Heroes Object
	heroes := []entity.Heroes{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT name, universe, skill, imageURL FROM heroes
	`)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Failed when retrive data heroes.",
		})
	}
	defer rows.Close()

	for rows.Next() {
		var h entity.Heroes

		err := rows.Scan(&h.Name, &h.Universe, &h.Skill, &h.ImageURL)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"msg": "Failed when retrive data heroes.",
			})
		}

		heroes = append(heroes, h)
	}

	json.NewEncoder(w).Encode(heroes)
}
