package config

import (
	"os"
	"strconv"
	"time"
)

// envString parses string environment variable.
func envString(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	return value
}

// envInt parses int environment variable.
func envInt(name string, defaultValue int) int {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

// envIntVariadic parses multiple int environment variables.
// Grabs first non-empty environment variable.
// If all of the provided environtment variables are empty, will use the defaultValue.
func envIntMultiple(names []string, defaultValue int) int {
	var value string
	for _, name := range names {
		value = os.Getenv(name)
		if value != "" {
			break
		}
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

// envDuration parses duration environment variable.
func envDuration(name string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}

// envStorageType parses StorageType environment variable.
func envStorageType(name string, defaultValue StorageType) StorageType {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	parsedValue := StorageType(value)
	if !parsedValue.IsValid() {
		return defaultValue
	}

	return parsedValue
}

// envBool parses bool environment variable.
func envBool(name string, defaultValue bool) bool {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}
