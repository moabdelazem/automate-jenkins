package main

import (
	"log"
	"net/http"

	"github.com/moabdelazem/automate-jenkins/internal/server"
)

func main() {
	router := server.NewRouter()

	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
