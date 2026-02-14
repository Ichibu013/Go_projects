package logic

import (
	"net/http"
	"restful-url-shortner/storage"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get the short code from URL path
	shortCode := r.URL.Path[1:] // slice '/' from URL

	originalUrl, found := storage.GetUrlMapping(shortCode)

	if !found {
		http.Error(w, "Short URl Not found", http.StatusNotFound)
		return
	}

	// 3. Redirect the user
	http.Redirect(w, r, originalUrl, http.StatusFound) // 302
}
