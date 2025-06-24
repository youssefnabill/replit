
package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Port string
	Env  string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		Port: getEnv("PORT", "5000"),
		Env:  getEnv("ENV", "development"),
	}

	return config, nil
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
