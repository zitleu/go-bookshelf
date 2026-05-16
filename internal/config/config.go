package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback

}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/bookshelf?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "jwtsecret"),
	}

}
