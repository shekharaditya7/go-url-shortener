package handlers

import (
	"net/http"              // For HTTP handling
	"url-shortener/storage" // DB functions
	"github.com/go-chi/chi/v5" // Router for URL parameters
)

// RedirectURL handles the redirection request
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// Extract the shortID from the URL
	shortID := chi.URLParam(r, "shortID")

	// Retrieve the original URL from the database
	var originalURL string
	err := storage.DB.QueryRow("SELECT original_url FROM urls WHERE short_id = ?", shortID).Scan(&originalURL)

	// Handle errors (shortID not found)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}
