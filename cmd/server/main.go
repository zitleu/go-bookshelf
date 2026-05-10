package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Bookshelf!")

	http.HandleFunc("/health", health)

	http.ListenAndServe(":8080", nil)

}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"status": "ok"}

	json.NewEncoder(w).Encode(response)
}
