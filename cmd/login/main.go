package main

import (
    "fmt"
    "log"
    "rest_api/internal/app/database"
)

func main() {
    // Mencoba mendapatkan koneksi ke database
    db, err := database.GetConnection()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer db.Close()

    // Test koneksi ke database
    if err := db.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    fmt.Println("Successfully connected to the database!")
}

