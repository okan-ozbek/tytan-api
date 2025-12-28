package config

import (
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port         string
	TimeoutRead  string
	TimeoutWrite string
	TimeoutIdle  string
	Debug        string
	Secret       string
}

func LoadConfig() *Config {
	return &Config{
		Database: *LoadDatabaseConfig(),
		Server:   *LoadServerConfig(),
	}
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   getEnv("DB_DRIVER", "sqlite", false),
		Host:     getEnv("DB_HOST", "", false),
		Port:     getEnv("DB_PORT", "", false),
		User:     getEnv("DB_USER", "", false),
		Password: getEnv("DB_PASSWORD", "", false),
		Name:     getEnv("DB_NAME", "", false),
	}
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         getEnv("SERVER_PORT", "8080", false),
		TimeoutRead:  getEnv("SERVER_TIMEOUT_READ", "15s", false),
		TimeoutWrite: getEnv("SERVER_TIMEOUT_WRITE", "15s", false),
		TimeoutIdle:  getEnv("SERVER_TIMEOUT_IDLE", "60s", false),
		Debug:        getEnv("SERVER_DEBUG", "false", false),
		Secret:       getEnv("SERVER_SECRET", "", true),
	}
}

// TODO revisit to add a better ENV library or methodology
func getEnv(key, defaultValue string, required bool) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if required {
		log.Fatalf("ENV VARIABLE: %s NOT SET!", key)
		return ""
	}

	log.Fatalf("ENV VARIABLE: %s, Using default: %s.", key, defaultValue)
	return defaultValue
}
