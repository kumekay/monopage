package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := initDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	r := mux.NewRouter()
	
	r.HandleFunc("/", createPageHandler).Methods("GET", "POST")
	r.HandleFunc("/{slug:[^_]+_edit}", editPageHandler).Methods("GET", "POST")
	r.HandleFunc("/{slug}", viewPageHandler).Methods("GET")
	
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}