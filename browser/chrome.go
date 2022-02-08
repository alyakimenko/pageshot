// Package browser contains browser related logic.
package browser

import (
	"context"

	"github.com/alyakimenko/pageshot/config"
	"github.com/chromedp/chromedp"
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
func (c *ChromeBrowser) Screenshot(ctx context.Context, url string) ([]byte, error) {
	allocCtx, cancel := c.allocateBrowser(ctx)
	defer cancel()

	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	var res []byte
	// quality is 100, will make it configurable later
	if err := chromedp.Run(ctx, c.fullpageScreenshot(url, &res, 100)); err != nil {
		return nil, err
	}

	return res, nil
}

// allocateBrowser allocates new chrome browser.
func (c *ChromeBrowser) allocateBrowser(ctx context.Context) (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.WindowSize(c.width, c.height),
	)

	return chromedp.NewExecAllocator(ctx, opts...)
}

// fullpageScreenshot takes a full screenshot with the specified image quality of the entire browser viewport.
func (c *ChromeBrowser) fullpageScreenshot(url string, res *[]byte, quility int) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(res, quility),
	}
}
