package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Ambil PORT dari env, fallback 4000
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Ambil DATABASE_DSN dari env
	dsn := os.Getenv("DATABASE_DSN")
	log.Printf("Using DATABASE_DSN: %s", dsn)

	// Health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port) // bind ke semua interface
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

