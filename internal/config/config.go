package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds configuration settings for the application.
type Config struct {
	Port string
}

// Load retrieves configuration from environment variables.
func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{Port: port}
}

// Address formats the port as an address string.
func (c *Config) Address() string {
	return ":" + c.Port
}
