package database

import (
    "database/sql"
    "fmt"
    "github.com/spf13/viper"
    _ "github.com/lib/pq"
    "time"
)

func GetConnection() (*sql.DB, error) {
    // Buat instance Viper
    config := viper.New()
    config.SetConfigName("config")
    config.SetConfigType("json")
    config.AddConfigPath("/storage/emulated/0/rest_api/configs")

    // Baca konfigurasi dari file JSON
    if err := config.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading config file: %w", err)
    }
    
    // Dapatkan nilai konfigurasi koneksi dari Viper
    host := config.GetString("database.host")
    port := config.GetInt("database.port")
    user := config.GetString("database.username") // Ubah sesuai dengan key yang sesuai
    password := config.GetString("database.password") 
    dbname := config.GetString("database.dbname")

    // Bentuk string koneksi ke PostgreSQL
    connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    // Buka koneksi ke PostgreSQL
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return nil, fmt.Errorf("error opening database connection: %w", err)
    }

    // Konfigurasi koneksi pool
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxIdleTime(5 * time.Minute)
    db.SetConnMaxLifetime(60 * time.Minute)

    return db, nil
}

