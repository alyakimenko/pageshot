package models

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
