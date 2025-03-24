package handlers

import (
	"encoding/json"           // For JSON encoding/decoding
	"math/rand"               // For random ID generation
	"net/http"                // For HTTP handling
	"url-shortener/storage"   // Import DB functions
)

// Request struct for incoming JSON payload
type ShortenRequest struct {
	URL string `json:"url"`
}

// Response struct for returning the short URL
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// Generates a random 6-character ID
func generateShortID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortID := make([]byte, 6)
	for i := range shortID {
		shortID[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortID)
}

// ShortenURL handles URL shortening requests
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	// Decode the JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a short ID
	shortID := generateShortID()

	// Store in the database
	_, err := storage.DB.Exec("INSERT INTO urls (short_id, original_url) VALUES (?, ?)", shortID, req.URL)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	// Send the short URL in response
	resp := ShortenResponse{
		ShortURL: "http://localhost:8080/" + shortID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
