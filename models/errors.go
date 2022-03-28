package models

import (
	"errors"
)

var (
	// ErrUnsupportedImageFormat occurs when unsupported or invalid image format was provided.
	ErrUnsupportedImageFormat = errors.New("unsupported image format")
)
