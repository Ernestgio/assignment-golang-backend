package config

import "os"

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getENV(key, defaultVal string) string {
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
		DBName:   getENV("DB_NAME", "bank"),
		Port:     getENV("DB_PORT", "5432"),
	}
)
