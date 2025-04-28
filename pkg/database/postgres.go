package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// InitDB menginisialisasi koneksi database dan mengembalikan *sqlx.DB
func ConnectDB() (*sqlx.DB, error) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "perpusapps"
	}
	
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "admin1234"
	}
	
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "manajemen_perpus"
	}

	// Format connection string yang benar untuk PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Buka koneksi dengan sqlx
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal terkoneksi ke PostgreSQL")
		return nil, err
	}

	// Test koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Gagal ping ke database")
		return nil, err
	}

	log.Info().Msg("Berhasil terkoneksi ke database")
	return db, nil
}