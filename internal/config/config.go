package config

// Config holds configuration for the speed engine
type Config struct {
	Port       string
	ServerName string
	Location   string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Port:       ":8080",
		ServerName: "Fluxmach Edge Node 1",
		Location:   "New York, US",
	}
}
