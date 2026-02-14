package logic

import (
	"encoding/json"
	"net/http"
	"restful-url-shortner/storage"
)

// ShortenHandler shortenHandler -> private method
// ShortenHandler -> public method
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Only allow Post Request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	// 2. Parse Incoming Data -> Make Request Body
	var mapping storage.URLMapping
	err := json.NewDecoder(r.Body).Decode(&mapping)
	if err != nil {
		http.Error(w, "Invalid JSON.", http.StatusBadRequest)
		return
	}

	// 3. Generate Short Code and Save
	shortCode := generateShortCode()
	storage.SaveUrlMapping(shortCode, mapping.OriginalURL)

	// 4. Return result as JSON -> Send Response Body
	response := map[string]string{
		"short_url":    "http://localhost:8080/" + shortCode,
		"original_url": mapping.OriginalURL,
	}

	// Set Header to tell to Browser this is JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
