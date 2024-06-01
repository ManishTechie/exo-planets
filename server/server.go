// Package server provides a thin wrapper around the default net/http Server with Gin Engine as its router.
package server

import (
	"context"
	"errors"
	"exo-planets/constants"
	"exo-planets/logging"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Start the HTTP server in a separate Go routine, and listen for OS kill signals.
// Whenever a signal is received, shut it down gracefully.
func Start(router http.Handler, serverName string, cleanupFns ...func() error) {

	logger := logging.GetLogger()

	// create the HTTP server
	srv := &http.Server{
		Addr:         ":" + os.Getenv(constants.PORT),
		Handler:      router,
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(5) * time.Second,
		IdleTimeout:  time.Duration(5) * time.Second,
	}

	// create the channel for the OS signal, used for shutdown
	stopCh := make(chan os.Signal, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP Server exited with a fatal error", zap.Error(err))
		}
	}()

	// notify the stopCh channel about OS interrupt, SIGINT and SIGTERM signals
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// block until a signal is received
	<-stopCh
	logger.Info("Shut down signal received")

	// create a context with timeout and cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cleanupErrors := []string{}
		for _, cleanupFn := range cleanupFns {
			if err := cleanupFn(); err != nil {
				cleanupErrors = append(cleanupErrors, err.Error())
			}
		}
		if len(cleanupErrors) > 0 {
			logger.Error("Failed to clean up properly after shutdown", zap.Error(errors.New(strings.Join(cleanupErrors, ", "))))
		}
		cancel()
	}()

	// initiate the graceful shutdown of the server within the previously created context
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Failed to gracefully shut down the Server", zap.Error(err))
	}
	logger.Info("Server shut down gracefully")
}
