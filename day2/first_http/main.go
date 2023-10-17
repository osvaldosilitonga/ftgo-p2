package main

import (
	"fmt"
	"net/http"
)

func handler() {

}

func main() {
	mux := http.NewServeMux()

	// Endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// Server struct
	app := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	// Listen and serve
	err := app.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
