package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bookshelf/monolith/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()
	addr := ":" + cfg.Port

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// routes
	r.Get("/health", health)

	fmt.Printf("Server starting on %s\n", addr)
	http.ListenAndServe(addr, r)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"status": "ok"}

	json.NewEncoder(w).Encode(response)
}
