package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// ScreenshotOptions is a various set of options
// that can be applied during the screenshotting process.
type ScreenshotOptions struct {
	// (required) URL of a page to take a screenshot of.
	URL string
	// (optional) Width of the browser viewport.
	Width int
	// (optional) Height of the browser viewport.
	Height int
	// (optional) Screenshot quality.
	Quality int
	// (optional) If true full page screenshot will be took.
	Fullpage bool
	// (optional) Viewport scale factor.
	Scale int
	// (optional) Delay before taking a screenshot.
	Delay int
	// (optional) Format of an image.
	Format ImageFormat
}

// Hash returns a hash of underlying ScreenshotOptions object.
func (so ScreenshotOptions) Hash() (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(fmt.Sprintf("%v", so)))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
