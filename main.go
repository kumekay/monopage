package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Single page routes
	r.HandleFunc("/", viewPageHandler).Methods("GET")
	r.HandleFunc("/edit", editPageHandler).Methods("GET", "POST")

	// Static file serving
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := host + ":" + port

	log.Printf("Server starting on %s", address)
	log.Fatal(http.ListenAndServe(address, r))
}
