package main

import (
	"log"                      // For logging messages
	"net/http"                  // For starting the HTTP server
	"url-shortener/handlers"    // Import handler functions
	"url-shortener/storage"     // Import database functions

	"github.com/go-chi/chi/v5"  // Chi router for routing
)

func main() {
	// Initialize the database (SQLite)
	storage.InitDB()

	// Create a new Chi router
	r := chi.NewRouter()

	// Define routes
	r.Post("/", handlers.ShortenURL)       // POST request → shorten URLs
	r.Get("/{shortID}", handlers.RedirectURL) // GET request → redirect to original URL

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
