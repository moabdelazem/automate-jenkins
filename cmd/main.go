package main

import (
	"log"
	"net/http"

	"github.com/moabdelazem/automate-jenkins/internal/server" // Adjust import path based on your go.mod module name
)

func main() {
	router := server.NewRouter()

	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err) // Use log.Fatalf for fatal errors
	}
}
