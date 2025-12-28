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
		Driver:   getEnv("DB_DRIVER", "sqlite"),
		Host:     getEnv("DB_HOST", ""),
		Port:     getEnv("DB_PORT", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", ""),
	}
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         getEnv("SERVER_PORT", "8080"),
		TimeoutRead:  getEnv("SERVER_TIMEOUT_READ", "15s"),
		TimeoutWrite: getEnv("SERVER_TIMEOUT_WRITE", "15s"),
		TimeoutIdle:  getEnv("SERVER_TIMEOUT_IDLE", "60s"),
		Debug:        getEnv("SERVER_DEBUG", "false"),
		Secret:       getEnv("SERVER_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Environment variable %s not set, using default value: %s", key, defaultValue)
	return defaultValue
}
