package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default kalau di lokal
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from Go server!")
    })

    log.Println("Server starting on port", port)
    err := http.ListenAndServe(":"+port, mux)
    if err != nil {
        log.Fatal(err)
    }
}

