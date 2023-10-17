package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Home Page")
}

func TestPage(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var person Person

	fmt.Fprintf(w, "Test params : %v", param.ByName("name"))
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(person)
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/test/:name", TestPage)

	log.Fatal(http.ListenAndServe(":8080", router))
}
