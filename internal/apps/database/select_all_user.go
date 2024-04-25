package database

import (
    "context"
    "fmt"
)

func CheckUser(name string) {
    // Mendapatkan koneksi ke database
    db, err := GetConnection() // Mengembalikan db dan err
    if err != nil {
        panic(err)
    }
    defer db.Close()


    ctx := context.Background()
    script := "SELECT username FROM users WHERE username = $1 LIMIT 1" // Menggunakan $1 sebagai placeholder
    fmt.Println(script)
    
    // Menyediakan nilai parameter ke dalam fungsi db.QueryContext()
    rows, err := db.QueryContext(ctx, script, name)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    if rows.Next() {
        var username string
        err := rows.Scan(&username)
        if err != nil {
            panic(err)
        }
        fmt.Println("Sukses Login", username)
    } else {
        fmt.Println("Gagal Login")
    }
}

