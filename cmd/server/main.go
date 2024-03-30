package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/andreis3/stores-ms/cmd/configs"
	"github.com/andreis3/stores-ms/cmd/proxy"
	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
)

func main() {
	mux := http.NewServeMux()
	log := logger.NewLogger()
	conf, err := configs.LoadConfig(".")
	if err != nil {
		log.Error(fmt.Sprintf("Error loading config: %s", err.Error()))
		os.Exit(1)
	}
	pool := postgres.NewPostgresDB(*conf)

	// HTTP server configuration
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", conf.ServerPort),
		Handler: mux,
	}

	// Start HTTP server in a goroutine
	go func() {
		proxy.ProxyDependency(mux, pool, log, conf)
		log.Info(fmt.Sprintf("Start server on port %s", conf.ServerPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Configure channel for shutdown signals
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Wait for shutdown signal
	sig := <-shutdownSignal

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initiate graceful shutdown
	log.Info(fmt.Sprintf(`Received signal: %s. Initiating graceful shutdown...`, strings.ToUpper(sig.String())))
	if err := server.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("Error during server shutdown: %s", err.Error()))
	}

	// Close connection pool with PostgreSQL
	pool.Close()
	close(shutdownSignal)
	log.Info("Completed close postgres connection...")
	log.Info("Shutdown complete exit code 0...")
}
