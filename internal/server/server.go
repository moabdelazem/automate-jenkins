package server

import (
	"github.com/gorilla/mux"
	"github.com/moabdelazem/automate-jenkins/internal/handlers"
)

// NewRouter creates and configures a new mux Router.
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/health", handlers.HealthHandler)

	return router
}
