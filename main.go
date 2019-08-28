package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting Server . . . ")
	r := mux.NewRouter()

	// Routes available
	r.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/analyze/logs", analyzeLogs).Methods(http.MethodPost)
	// starting server
	log.Fatal(http.ListenAndServe(":8000", r))
}
