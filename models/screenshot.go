package models

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
}
