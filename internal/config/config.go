package config

import "os"

// Config holds configuration for the speed engine
type Config struct {
	Port       string
	ServerName string
	Location   string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "Fluxmach Edge Node 1"
	}

	location := os.Getenv("LOCATION")
	if location == "" {
		location = "New York, US"
	}

	return &Config{
		Port:       ":8080",
		ServerName: serverName,
		Location:   location,
	}
}
