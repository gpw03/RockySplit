package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
	Time string `json:"time"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Content-Type", "application/json")

	resp := HealthResponse{
		Status: "ok",
		Time: time.Now().Format(time.RFC3339),
	}

	// Encod the struct as JSON and write to response w
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// If encoding goes wrong return a 500 error
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/healthz", healthHandler)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
