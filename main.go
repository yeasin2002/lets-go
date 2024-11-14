package main

import (
	"fmt"
	"log"
	"net/http"
)

// homeHandler is a simple HTTP handler that responds with a welcome message.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Welcome to the Go HTTP API!"); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
	}
}

// main function initializes the HTTP server and routes.
func main() {
	// Define routes
	http.HandleFunc("/", homeHandler)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
