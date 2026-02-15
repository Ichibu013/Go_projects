package logic

import (
	"encoding/json"
	"log"
	"net/http"
	"restful-url-shortner/storage"
)

type Handler struct {
	Repo   *storage.Repository
	Logger *log.Logger
}

// ShortenHandler shortenHandler -> private method
// ShortenHandler -> public method
func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Only allow Post Request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	// 2. Parse Incoming Data -> Make Request Body
	var mapping struct {
		OriginalURL string `json:"original_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&mapping); err != nil {
		h.Logger.Println("Invalid Request")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 3. Generate Short Code and Save
	shortCode := generateShortCode()

	if err := h.Repo.SaveURL(shortCode, mapping.OriginalURL); err != nil {
		h.Logger.Println("Error Saving URL to DB: %v", err)
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	// 4. Return result as JSON -> Send Response Body
	response := map[string]string{
		"short_url":    "http://localhost:8080/" + shortCode,
		"original_url": mapping.OriginalURL,
	}

	// Set Header to tell to Browser this is JSON
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
