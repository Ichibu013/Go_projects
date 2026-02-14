package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"restful-url-shortner/logic"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// w: The Response Writer (write Data)
	// r: Request
	_, err := fmt.Fprint(w, "Welcome to url Shortener")
	if err != nil {
		return
	}
}

// Manually Create a server to listen on port 8080
func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Register the route
	http.HandleFunc("/shorten", logic.ShortenHandler)

	http.HandleFunc("/", logic.RedirectHandler)

	// 2. Start Server on port 8080
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
