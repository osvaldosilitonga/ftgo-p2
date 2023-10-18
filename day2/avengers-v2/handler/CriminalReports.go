package handler

import (
	"avengers-v2/entity"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
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

	// Criminal Reports Object
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

func (handler Criminal) GetCriminalReportById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer handler.DB.Close()

	w.Header().Set("content-type", "application/json")

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Invalid Param ID",
		})
		return
	}

	// Criminal Reports Object
	criminal := entity.CriminalReports{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = handler.DB.QueryRowContext(ctx, `
	SELECT CriminalReports.id, heroes.name, villain.name, location, date, description, status
	FROM CriminalReports
	JOIN heroes ON CriminalReports.hero_id = heroes.id
	JOIN villain ON CriminalReports.villain_id = villain.id
	WHERE CriminalReports.id = ?
	`, id).Scan(&criminal.ID, &criminal.Hero, &criminal.Villain, &criminal.Location, &criminal.Date, &criminal.Description, &criminal.Status)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Failed when retrive data criminal.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(criminal)
}

func (handler Criminal) PostCriminalReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer handler.DB.Close()

	w.Header().Set("content-type", "application/json")

	var criminal entity.Criminal

	err := json.NewDecoder(r.Body).Decode(&criminal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Insert Failed",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = handler.DB.ExecContext(ctx, `
	INSERT INTO CriminalReports (hero_id, villain_id, location, date, description, status)
	VALUES (?,?,?,?,?,?)
	`, criminal.HeroId, criminal.VillainId, criminal.Location, criminal.Date, criminal.Description, criminal.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Insert Failed",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Insert Success",
	})
}

func (handler Criminal) PutCriminalReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer handler.DB.Close()

	w.Header().Set("content-type", "application/json")

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Invalid Param ID",
		})
		return
	}

	criminal := entity.Criminal{}

	// body data
	err = json.NewDecoder(r.Body).Decode(&criminal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Update Failed",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := handler.DB.ExecContext(ctx, `
	UPDATE CriminalReports
	SET hero_id = ?, villain_id = ?, location = ?, date = ?, description = ?, status = ?
	WHERE id = ?
	`, criminal.HeroId, criminal.VillainId, criminal.Location, criminal.Date, criminal.Description, criminal.Status, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Upadate Failed.",
		})
		return
	}

	aff, err := res.RowsAffected()
	if aff == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Update Failed.",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Update Success",
	})
}
