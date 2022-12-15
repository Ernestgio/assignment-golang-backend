package config

import (
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
	godotenv.Load()
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var (
	ENV       = getENV("ENV", "testing")
	AppName   = "assignment-golang-backend"
	SecretKey = getENV("SECRET_KEY", "secret")
	DBConfig  = dbConfig{
		Host:     getENV("DB_HOST", "localhost"),
		User:     getENV("DB_USER", "postgres"),
		Password: getENV("DB_PASSWORD", "postgres"),
		DBName:   getENV("DB_NAME", "wallet_db_ernest"),
		Port:     getENV("DB_PORT", "5432"),
	}
)
