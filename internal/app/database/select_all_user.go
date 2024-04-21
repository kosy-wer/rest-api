package database

import (
    "context"
    "database/sql"
    "fmt"
)

func CheckUser() {
    // Mendapatkan koneksi ke database
    db, err := GetConnection() // Mengembalikan db dan err
    if err != nil {
        panic(err)
    }
    defer db.Close()

    ctx := context.Background()

    script := "SELECT user_id, username, password, email FROM users"
    rows, err := db.QueryContext(ctx, script)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var username, password string
        var user_id int32
        var email sql.NullString

        err = rows.Scan(&user_id, &username, &password, &email)
        if err != nil {
            panic(err)
        }
        fmt.Println("================")
        fmt.Println("Id:", user_id)
        fmt.Println("Name:", username)
        if email.Valid {
            fmt.Println("Email:", email.String)
        } else {
            fmt.Println("Email: (empty)")
        }
        fmt.Println("Password:", password)
    }
}

