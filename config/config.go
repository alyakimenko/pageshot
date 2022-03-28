// Package config holds all config related things.
package config

import (
	"os"
	"time"
)

// Config holds all configurable values for the app.
type Config struct {
	Server  ServerConfig
	Browser BrowserConfig
	Storage StorageConfig
	Logger  LoggerConfig
}

// NewConfig initializes new Config based on the environmental variables.
func NewConfig() Config {
	return Config{
		Server: ServerConfig{
			Port:         envInt("SERVER_PORT", 8000),
			ReadTimeout:  envDuration("SERVER_READ_TIMEOUT", 5*time.Second),
			WriteTimeout: envDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  envDuration("SERVER_IDLE_TIMEOUT", 5*time.Second),
		},
		Browser: BrowserConfig{
			Width:  envInt("BROWSER_WIDTH", 1440),
			Height: envInt("BROWSER_HEIGHT", 900),
		},
		Storage: StorageConfig{
			Type:      envStorageType("STORAGE_TYPE", "local"),
			Directory: envString("STORAGE_DIRECTORY", os.TempDir()),
		},
		Logger: LoggerConfig{
			Level: envString("LOGGER_LEVEL", "INFO"),
		},
	}
}
