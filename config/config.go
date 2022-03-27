// Package config holds all config related things.
package config

import (
	"fmt"
	"time"
)

// Config holds all configurable values for the app.
type Config struct {
	Server  ServerConfig
	Browser BrowserConfig
	Logger  LoggerConfig
}

// ServerConfig holds HTTP server related configurable values.
type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// Addr returns server address in the format [host]:[port].
func (sc ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", sc.Port)
}

// BrowserConfig holds browser related configurable values.
type BrowserConfig struct {
	Width  int
	Height int
}

// LoggerConfig holds logger related configurable values.
type LoggerConfig struct {
	Level string `default:"INFO"`
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
		Logger: LoggerConfig{
			Level: envString("LOGGER_LEVEL", "INFO"),
		},
	}
}
