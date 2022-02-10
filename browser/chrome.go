// Package browser contains browser related logic.
package browser

import (
	"context"

	"github.com/alyakimenko/pageshot/config"
	"github.com/alyakimenko/pageshot/models"
	"github.com/chromedp/chromedp"
)

const (
	// maxQuality is a maximum allowed quality for screenshots.
	maxQuality = 100

	// defaultQuality is a default quality level for screenshots.
	// This value will be used if a custom value is not provided.
	defaultQuality = 70
)

// ChromeBrowser operates Chrome browser logic.
type ChromeBrowser struct {
	width  int
	height int
}

// ChromeBrowserParams is an incoming params for the NewChromeBrowser function.
type ChromeBrowserParams struct {
	Config config.BrowserConfig
}

// NewChromeBrowser creates new instance of the ChromeBrowser.
func NewChromeBrowser(params ChromeBrowserParams) *ChromeBrowser {
	return &ChromeBrowser{
		width:  params.Config.Width,
		height: params.Config.Height,
	}
}

// Screenshot takes a screenshot based on the provided parameters.
func (c *ChromeBrowser) Screenshot(ctx context.Context, opts models.ScreenshotOptions) ([]byte, string, error) {
	allocCtx, cancel := c.allocateBrowser(ctx)
	defer cancel()

	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	var res []byte
	if err := chromedp.Run(ctx, c.fullpageScreenshot(&res, opts)); err != nil {
		return nil, "", err
	}

	contentType := "image/jpg"
	if opts.Quality == maxQuality {
		contentType = "image/png"
	}

	return res, contentType, nil
}

// allocateBrowser allocates new chrome browser.
func (c *ChromeBrowser) allocateBrowser(ctx context.Context) (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.WindowSize(c.width, c.height),
	)

	return chromedp.NewExecAllocator(ctx, opts...)
}

// fullpageScreenshot takes a full screenshot with the specified image quality of the entire browser viewport.
func (c *ChromeBrowser) fullpageScreenshot(res *[]byte, opts models.ScreenshotOptions) chromedp.Tasks {
	var tasks chromedp.Tasks
	if opts.Width != 0 && opts.Height != 0 {
		tasks = append(tasks, chromedp.EmulateViewport(int64(opts.Width), int64(opts.Height)))
	}

	tasks = append(tasks, chromedp.Navigate(opts.URL))
	if opts.Quality == 0 {
		opts.Quality = defaultQuality
	}

	if opts.Fullpage {
		tasks = append(tasks, chromedp.FullScreenshot(res, opts.Quality))
	} else {
		tasks = append(tasks, chromedp.CaptureScreenshot(res))
	}

	return tasks
}
