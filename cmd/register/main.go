package main

import (
    "rest_api/internal/apps/database"
)

func main() {
    // Mencoba mendapatkan koneksi ke database
 database.CheckUser("joko")
}

