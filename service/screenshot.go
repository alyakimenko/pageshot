package service

import "context"

// Browser is a browser interface that is required for the ScreenshotService.
type Browser interface {
	Screenshot(ctx context.Context, url string) ([]byte, error)
}

// ScreenshotService deals with screenshots related logic.
type ScreenshotService struct {
	browser Browser
}

// ScreenshotServiceParams is an incoming params for the NewScreenshotService function.
type ScreenshotServiceParams struct {
	Browser Browser
}

// NewScreenshotService creates new instance of the ScreenshotService.
func NewScreenshotService(params ScreenshotServiceParams) *ScreenshotService {
	return &ScreenshotService{
		browser: params.Browser,
	}
}

// Screenshot takes a screenshot via the underlying Browser.
func (s *ScreenshotService) Screenshot(ctx context.Context, url string) ([]byte, error) {
	return s.browser.Screenshot(ctx, url)
}
