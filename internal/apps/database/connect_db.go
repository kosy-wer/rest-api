package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
	//"github.com/spf13/viper"
)

func GetConnection() (*sql.DB, error) {
	config := viper.New()
	/*config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("/data/data/com.termux/files/home/go/configs") // path configs

	// Baca file konfigurasi
	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
 */
	// Ambil DSN
	//dsn := config.GetString("database.dsn")
	dsn := os.Getenv("DATABASE_DSN")

	// Koneksi ke database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Setup pool
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}

