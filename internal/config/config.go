package config

// Config holds configuration for the speed engine
type Config struct {
	Port string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Port: ":8080",
	}
}
