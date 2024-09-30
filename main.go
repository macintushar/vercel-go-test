package main

import (
	"log"
	"net/http"

	"github.com/macintushar/vercel-go-test/api"
)

func main() {
	http.HandleFunc("/api", api.Handler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
