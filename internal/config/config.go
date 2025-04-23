package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL	string
	Port	string
	JWTSecret	string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		DBURL:	fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
		Port:	os.Getenv("APP_PORt"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}