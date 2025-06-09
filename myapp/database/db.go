package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Ambil konfigurasi dari environment
	dbURL := os.Getenv("DATABASE_URL") // e.g. postgres://user:123@db:5432/summary_db?sslmode=disable

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Cek koneksi
	if err = DB.Ping(); err != nil {
		log.Fatalf("Gagal ping database: %v", err)
	}

	// Buat tabel jika belum ada
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS summaries (
        id SERIAL PRIMARY KEY,
        input TEXT NOT NULL,
        summary TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW()
    )`)
	if err != nil {
		log.Fatalf("Gagal buat tabel: %v", err)
	}
}
