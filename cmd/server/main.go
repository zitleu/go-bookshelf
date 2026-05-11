package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bookshelf/monolith/internal/config"
)

func main() {
	cfg := config.Load()

	http.HandleFunc("/health", health)

	addr := ":" + cfg.Port
	fmt.Printf("Server starting on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"status": "ok"}

	json.NewEncoder(w).Encode(response)
}
