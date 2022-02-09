// Package config holds all config related things.
package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config holds all configurable values for the app.
type Config struct {
	Server  ServerConfig  `split_words:"true"`
	Browser BrowserConfig `split_words:"true"`
	Logger  LoggerConfig  `split_words:"true"`
}

// ServerConfig holds HTTP server related configurable values.
type ServerConfig struct {
	Host         string        `default:"localhost"`
	Port         string        `default:"8000"`
	ReadTimeout  time.Duration `default:"10s" split_words:"true"`
	WriteTimeout time.Duration `default:"5s" split_words:"true"`
	IdleTimeout  time.Duration `default:"5s" split_words:"true"`
}

// Addr returns server address in the format [host]:[port].
func (sc ServerConfig) Addr() string {
	return ":8000"
}

// BrowserConfig holds browser related configurable values.
type BrowserConfig struct {
	Width  int `default:"1440"`
	Height int `default:"900"`
}

// LoggerConfig holds logger related configurable values.
type LoggerConfig struct {
	Level string `default:"INFO"`
}

// NewConfig initializes new Config based on the environmental variables.
func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
