package models

// ScreenshotOptions is a various set of options
// that can be applied during the screenshotting process.
type ScreenshotOptions struct {
	URL      string
	Width    int
	Height   int
	Quality  int
	Fullpage bool
}
