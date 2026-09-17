package config

import(
	"os"
)

type Config struct {
	Port 		string
	DatabaseURL string
	JWTSecret	string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")

	return &Config{
		Port:	port,
		DatabaseURL: dbURL,
	}
}