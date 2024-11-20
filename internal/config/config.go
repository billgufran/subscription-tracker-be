package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
	URL      string // For Railway's DATABASE_URL
}

type JWTConfig struct {
	SecretKey       string
	ExpirationHours int
}

// Load initializes configuration from environment variables
func Load() *Config {
	config := &Config{
		Server: ServerConfig{
			Port: getEnvOrDefault("PORT", "8080"), // Railway uses PORT
		},
		Database: loadDatabaseConfig(),
		JWT: JWTConfig{
			SecretKey:       getEnvOrDefault("JWT_SECRET_KEY", "your-secret-key"),
			ExpirationHours: getEnvAsIntOrDefault("JWT_EXPIRATION_HOURS", 24),
		},
	}

	return config
}

func loadDatabaseConfig() DatabaseConfig {
	// Railway provides DATABASE_URL
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return DatabaseConfig{
			URL: dbURL,
		}
	}

	// Fallback to individual connection parameters
	return DatabaseConfig{
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
		Name:     getEnvOrDefault("DB_NAME", "subscription_tracker"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		SSLMode:  getEnvOrDefault("DB_SSL_MODE", "disable"),
	}
}

// DatabaseConfig GetDSN returns the Data Source Name (DSN) for the database.
func (c *DatabaseConfig) GetDSN() string {
	if c.URL != "" {
		return c.URL
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.Name, c.Port, c.SSLMode,
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Warning: Invalid integer value for %s, using default: %d", key, defaultValue)
	}
	return defaultValue
}
