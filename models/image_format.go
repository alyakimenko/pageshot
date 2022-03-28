package models

import (
	"strings"
)

// ImageFormat represents an image format.
type ImageFormat string

const (
	// PNGImageFormat represents PNG image format.
	PNGImageFormat ImageFormat = "png"
	// JPEGImageFormat represents JPEG image format.
	JPEGImageFormat ImageFormat = "jpeg"
	// WebPImageFormat represents WebP image format.
	WebPImageFormat ImageFormat = "webp"
)

// IsValid validates the ImageFormat.
func (f ImageFormat) IsValid() bool {
	return f == PNGImageFormat || f == JPEGImageFormat || f == WebPImageFormat
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
// It's mainly for use within gorilla/schema decoding.
func (f *ImageFormat) UnmarshalText(text []byte) (err error) {
	*f, err = ParseImageFormat(string(text))

	return
}

// ContentType returns image format as a Content-Type header value.
func (f ImageFormat) ContentType() string {
	switch f {
	case PNGImageFormat:
		return "image/png"
	case JPEGImageFormat:
		return "image/jpeg"
	case WebPImageFormat:
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}

// Extension returns corresponding file extension for the underlying ImageFormat.
func (f ImageFormat) Extension() string {
	switch f {
	case PNGImageFormat:
		return ".png"
	case JPEGImageFormat:
		return ".jpeg"
	case WebPImageFormat:
		return ".webp"
	default:
		return ""
	}
}

// ParseImageFormat parses the ImageFormat.
func ParseImageFormat(text string) (ImageFormat, error) {
	switch strings.ToLower(text) {
	case "png":
		return PNGImageFormat, nil
	case "jpeg":
		return JPEGImageFormat, nil
	case "webp":
		return WebPImageFormat, nil
	default:
		return "", ErrUnsupportedImageFormat
	}
}
