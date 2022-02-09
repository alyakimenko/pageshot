package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alyakimenko/pageshot/browser"
	"github.com/alyakimenko/pageshot/config"
	"github.com/alyakimenko/pageshot/logger"
	"github.com/alyakimenko/pageshot/server"
	v1 "github.com/alyakimenko/pageshot/server/v1"
	"github.com/alyakimenko/pageshot/service"
)

func main() {
	// init global config
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// init logger
	logger, err := logger.NewLogrusLogger(config.Logger)
	if err != nil {
		panic(err)
	}

	// init browser
	chromeBrowser := browser.NewChromeBrowser(browser.ChromeBrowserParams{
		Config: config.Browser,
	})

	// create screenshot service with the browser
	screenshotService := service.NewScreenshotService(service.ScreenshotServiceParams{
		Browser: chromeBrowser,
	})

	// init v1 HTTP handler
	handler := v1.NewHandler(v1.HandlerParams{
		Logger:            logger,
		ScreenshotService: screenshotService,
	})

	// create HTTP server with the initialized v1 handler
	server := server.NewServer(server.Params{
		Config:  config.Server,
		Handler: handler,
	})

	// start the server
	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("server caught an error: %s\n", err.Error())
		}
	}()

	// graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		logger.Errorf("failed to stop server: %s\n", err.Error())
	}
}
