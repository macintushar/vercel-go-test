package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello from Vercel Edge Functions!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
