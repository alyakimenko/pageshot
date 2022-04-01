// Package browser contains browser related logic.
package browser

import (
	"context"
	"time"

	"github.com/alyakimenko/pageshot/config"
	"github.com/alyakimenko/pageshot/models"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

const (
	// maxQuality is a maximum allowed quality level for screenshots.
	maxQuality = 100

	// defaultQuality is a default quality level for screenshots.
	// This value will be used if a custom value is not provided.
	defaultQuality = 70
)

// ChromeBrowser operates Chrome browser logic.
type ChromeBrowser struct {
	width  int
	height int
	url    string
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
		url:    params.Config.URL,
	}
}

// Screenshot takes a screenshot based on the provided parameters.
func (c *ChromeBrowser) Screenshot(ctx context.Context, opts models.ScreenshotOptions) ([]byte, error) {
	var allocCtx context.Context
	var cancel context.CancelFunc

	// try to connect to a remote browser, otherwise allocate a local one
	if c.url != "" {
		allocCtx, cancel = c.allocateRemoteBrowser(ctx, c.url)
	} else {
		allocCtx, cancel = c.allocateLocalBrowser(ctx)
	}
	defer cancel()

	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	return c.screenshot(ctx, opts)
}

// allocateBrowser allocates new local chrome browser.
func (c *ChromeBrowser) allocateLocalBrowser(ctx context.Context) (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.WindowSize(c.width, c.height),
	)

	return chromedp.NewExecAllocator(ctx, opts...)
}

// connectRemoteBrowser connects to a remote browser by the provided url.
// The url should point to the browser's websocket address.
func (c *ChromeBrowser) allocateRemoteBrowser(ctx context.Context, url string) (context.Context, context.CancelFunc) {
	return chromedp.NewRemoteAllocator(ctx, url)
}

// fullpageScreenshot takes a screenshot with the specified screenshot options.
func (c *ChromeBrowser) screenshot(ctx context.Context, opts models.ScreenshotOptions) ([]byte, error) {
	var tasks chromedp.Tasks

	if opts.Width == 0 {
		opts.Width = c.width
	}

	if opts.Height == 0 {
		opts.Height = c.height
	}

	tasks = append(tasks,
		chromedp.EmulateViewport(
			int64(opts.Width), int64(opts.Height),
			chromedp.EmulateScale(float64(opts.Scale)/100),
		),
	)

	tasks = append(tasks, chromedp.Navigate(opts.URL))

	if opts.Quality == 0 || opts.Quality > maxQuality {
		opts.Quality = defaultQuality
	}

	if opts.Delay != 0 {
		tasks = append(tasks, chromedp.Sleep(time.Duration(opts.Delay)*time.Millisecond))
	}

	var res []byte
	if opts.Fullpage {
		tasks = append(tasks, c.fullpageScreenshot(&res, opts))
	} else {
		tasks = append(tasks, c.captureScreenshot(&res, opts))
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		return nil, err
	}

	return res, nil
}

// fullpageScreenshot takes a screenshot of an entire page.
func (c *ChromeBrowser) fullpageScreenshot(res *[]byte, opts models.ScreenshotOptions) chromedp.EmulateAction {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		// get layout metrics
		_, _, contentSize, _, _, cssContentSize, err := page.GetLayoutMetrics().Do(ctx)
		if err != nil {
			return err
		}
		// protocol v90 changed the return parameter name (contentSize -> cssContentSize)
		if cssContentSize != nil {
			contentSize = cssContentSize
		}

		format := page.CaptureScreenshotFormat(opts.Format)

		// capture screenshot
		*res, err = page.CaptureScreenshot().
			WithCaptureBeyondViewport(true).
			WithFormat(format).
			WithQuality(int64(opts.Quality)).
			WithClip(&page.Viewport{
				X:      0,
				Y:      0,
				Width:  contentSize.Width,
				Height: contentSize.Height,
				Scale:  1,
			}).Do(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}

// captureScreenshot takes a screenshot of a current viewport.
func (c *ChromeBrowser) captureScreenshot(res *[]byte, opts models.ScreenshotOptions) chromedp.EmulateAction {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		format := page.CaptureScreenshotFormat(opts.Format)

		var err error
		*res, err = page.CaptureScreenshot().
			WithFormat(format).
			WithCaptureBeyondViewport(true).
			Do(ctx)

		return err
	})
}
