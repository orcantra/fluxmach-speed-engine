package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds configuration for the speed engine
type Config struct {
	Port       string
	ServerName string
	Location   string
	AuthKey    string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "Fluxmach Edge Node 1"
	}

	location := os.Getenv("LOCATION")
	if location == "" {
		location = "New York, US"
	}

	authKey := os.Getenv("FLUXMACH_AUTH_KEY")
	if authKey == "" {
		// Log a warning if AuthKey is missing in a production-like environment
		log.Println("WARNING: FLUXMACH_AUTH_KEY is not set!")
	}

	return &Config{
		Port:       ":8080",
		ServerName: serverName,
		Location:   location,
		AuthKey:    authKey,
	}
}
