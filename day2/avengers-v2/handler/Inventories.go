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

type Inventory struct {
	DB *sql.DB
}

func NewInventoryHandler(db *sql.DB) Inventory {
	return Inventory{
		DB: db,
	}
}

func (handler Inventory) GetInventories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	// Inventories Object
	inventories := []entity.Inventories{}
	// var inventories []entity.Inventories

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := handler.DB.QueryContext(ctx, `
	SELECT name, stock FROM inventories
	`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Failed when retrive data inventories.",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i entity.Inventories

		err := rows.Scan(&i.Name, &i.Stock)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"msg": "Failed when retrive data inventories.",
			})
			return
		}

		inventories = append(inventories, i)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventories)
}

func (handler Inventory) GetInventoriesById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Invalid Param ID",
		})
		return
	}

	w.Header().Set("content-type", "application/json")

	// Inventories Object
	inventories := entity.Inventories{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = handler.DB.QueryRowContext(ctx, `
	SELECT name, stock FROM inventories
	WHERE id = ?
	`, id).Scan(&inventories.Name, &inventories.Stock)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Failed when retrive data inventories.",
		})
		return
	}

	json.NewEncoder(w).Encode(inventories)
}

func (handler Inventory) PostInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	var inventory entity.Inventories

	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = handler.DB.ExecContext(ctx, `
	INSERT INTO inventories (name, stock)
	VALUES (?,?)
	`, inventory.Name, inventory.Stock)
	if err != nil {
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

func (handler Inventory) PutInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Invalid Param ID",
		})
		return
	}

	var inventory entity.Inventories

	// body data
	err = json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := handler.DB.ExecContext(ctx, `
	UPDATE inventories
	SET name = ?, stock = ?
	WHERE id = ?
	`, inventory.Name, inventory.Stock, id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Upadate Failed.",
		})
		return
	}

	aff, err := res.RowsAffected()
	if aff == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Update Failed.",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Update Success",
	})
}

func (handler Inventory) DeleteInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Invalid Param ID",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := handler.DB.ExecContext(ctx, `
	DELETE FROM inventories
	WHERE id = ?
	`, id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Delete Failed.",
		})
		return
	}

	aff, err := res.RowsAffected()
	if aff == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Delete Failed.",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Delete Success",
	})
}
