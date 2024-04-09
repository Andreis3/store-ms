package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/util"
)

func GracefulShutdown(server *http.Server, pool *postgres.Postgres, log *logger.Logger) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-shutdownSignal
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info("Initiating graceful shutdown...")
	if err := server.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("Error during server shutdown: %s", err.Error()))
	}
	log.Info("Closing postgres connection...")
	pool.Close()
	log.Info("Shutdown complete exit code 0...")
	os.Exit(util.EXIT_SUCCESS)
}
