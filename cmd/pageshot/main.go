package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alyakimenko/pageshot/browser"
	"github.com/alyakimenko/pageshot/config"
	"github.com/alyakimenko/pageshot/logger"
	"github.com/alyakimenko/pageshot/service"
	"github.com/alyakimenko/pageshot/storage/local"
	"github.com/alyakimenko/pageshot/transport/rest"
)

func main() {
	// init global config
	config := config.NewConfig()

	// init logger
	logger, err := logger.NewLogrusLogger(config.Logger)
	if err != nil {
		panic(err)
	}

	// init browser
	chromeBrowser := browser.NewChromeBrowser(browser.ChromeBrowserParams{
		Config: config.Browser,
	})

	// init storage
	localStorage := local.NewStorage(local.StorageParams{
		Config: config.Storage,
	})

	// create screenshot service with the browser
	screenshotService := service.NewScreenshotService(service.ScreenshotServiceParams{
		Browser:     chromeBrowser,
		FileStorage: localStorage,
	})

	// init v1 HTTP handler
	handler := rest.NewHandler(rest.HandlerParams{
		Logger:            logger,
		ScreenshotService: screenshotService,
	})

	// create HTTP server with the initialized v1 handler
	server := rest.NewServer(rest.ServerParams{
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

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Errorf("failed to stop server: %s\n", err.Error())
	}
}
