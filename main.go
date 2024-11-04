package main

import (
	"log"
	"net/http"

	api "github.com/ashjshaw/go_kv_database/internal/app"
	"github.com/gorilla/mux"
)

func main() {
	handler := api.New()
	router := mux.NewRouter()
	router.HandleFunc("/{key}", handler.GetHandler).Methods("GET")
	router.HandleFunc("/{key}", handler.PutHandler).Methods("PUT")
	router.HandleFunc("/{key}", handler.DeleteHandler).Methods("DELETE")
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
