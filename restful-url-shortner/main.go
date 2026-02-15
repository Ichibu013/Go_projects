package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"restful-url-shortner/logic"
	"restful-url-shortner/storage"
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

	constStr := os.Getenv("DB_STRING")
	if constStr == "" {
		constStr = "postgres://postgres:root@172.25.8.167:5432/postgres?sslmode=disable"
	}

	repo, err := storage.NewRepository(constStr)
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

	// Dependency Injection
	h := &logic.Handler{
		Repo:   repo,
		Logger: logger,
	}

	// 1. Register the route
	http.HandleFunc("/shorten", h.ShortenHandler)

	http.HandleFunc("/", h.RedirectHandler)

	// 2. Start Server on port 8080
	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
