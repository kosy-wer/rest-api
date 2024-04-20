package main

import (
    "fmt"
    "context"
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

    ctx := context.Background()

	script := "INSERT INTO users(username,password,email) VALUES('joko', 'password678','joko@gmail.com')"
	_, err = db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")


    // Test koneksi ke database
    if err := db.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }


    fmt.Println("Successfully connected to the database!")
}

