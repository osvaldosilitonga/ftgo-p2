package handler

import (
	"avengers-v2/entity"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Criminal struct {
	DB *sql.DB
}

func NewCriminalHandler(db *sql.DB) Criminal {
	return Criminal{
		DB: db,
	}
}

func (handler Criminal) GetCriminalReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer handler.DB.Close()
	w.Header().Set("content-type", "application/json")

	// Inventories Object
	criminal := []entity.CriminalReports{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := handler.DB.QueryContext(ctx, `
	SELECT CriminalReports.id, heroes.name, villain.name, location, date, description, status
	FROM CriminalReports
	JOIN heroes ON CriminalReports.hero_id = heroes.id
	JOIN villain ON CriminalReports.villain_id = villain.id;
	`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Failed when retrive data criminal reports.",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cr entity.CriminalReports

		err := rows.Scan(&cr.ID, &cr.Hero, &cr.Villain, &cr.Location, &cr.Date, &cr.Description, &cr.Status)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"msg": "Failed when retrive data criminal reports.",
			})
			return
		}

		criminal = append(criminal, cr)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(criminal)
}
