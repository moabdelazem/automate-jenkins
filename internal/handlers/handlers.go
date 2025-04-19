package handlers

import (
	"encoding/json"
	"log" // Import the log package
	"net/http"
)

// WriteJSON is a helper function for writing JSON responses.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return // Don't encode nil data
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Log the error internally, but send a generic message to the client
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// RootHandler handles requests to the root path "/".
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: Method=%s Path=%s RemoteAddr=%s UserAgent=%s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent()) // Add more details
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

// HealthHandler handles requests to the "/health" path.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: Method=%s Path=%s RemoteAddr=%s UserAgent=%s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent()) // Add more details
	response := map[string]string{"status": "ok"}
	WriteJSON(w, http.StatusOK, response)
}
