package main

import (
	"avengers2/config"
	"avengers2/entity"
	"avengers2/hanlder"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// type Heroes struct {
// 	Name string
// 	Universe string
// 	Skill string
// 	ImageURL string
// }

func GetHeroes() []entity.Heroes {
	db, err := config.DBConn()
	if err != nil {
		log.Fatal("Can't connect to DB")
	}
	defer db.Close()

	heroes, err := hanlder.GetHeroes(db)
	if err != nil {
		log.Fatal("Failed to retrive heroes data from db, [err] :", err)
	}

	return heroes
}

func GetVillain() []entity.Villain {
	db, err := config.DBConn()
	if err != nil {
		log.Fatal("Can't connect to DB")
	}
	defer db.Close()

	villain, err := hanlder.GetVillain(db)
	if err != nil {
		log.Fatal("Failed to retrive heroes data from db, [err] :", err)
	}

	return villain
}

func main() {
	// Handler
	mux := http.NewServeMux()
	// Endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Avengers")
	})

	mux.HandleFunc("/heroes", func(w http.ResponseWriter, r *http.Request) {
		heroes := GetHeroes()
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(heroes)
	})

	mux.HandleFunc("/villain", func(w http.ResponseWriter, r *http.Request) {
		villain := GetVillain()
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(villain)
	})

	// Server
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server not connect")
	}
}
