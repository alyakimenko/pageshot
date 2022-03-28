package config

import (
	"fmt"
	"time"
)

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
