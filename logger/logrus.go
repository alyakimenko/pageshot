package logger

import (
	"fmt"

	"github.com/alyakimenko/pageshot/config"
	"github.com/sirupsen/logrus"
)

// NewLogrusLogger initizalies new *logrus.Logger.
func NewLogrusLogger(config config.LoggerConfig) (*logrus.Logger, error) {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}

	logger.SetLevel(level)

	return logger, nil
}
