package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
	go func() {
		proxy.ProxyDependency(mux, pool, log, conf)
		log.Info(fmt.Sprintf("Start server on port %s", conf.ServerPort))
		if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), mux); err != nil {
			log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
			os.Exit(1)
		}
	}()
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-shutdownSignal
	defer func() {
		log.Info(fmt.Sprintf(`Received signal: %s. Initiating graceful shutdown...`, strings.ToUpper(sig.String())))
		pool.Close()
		log.Info("Completed close postgres connection...")
		log.Info("Shutdown complete exit code 0...")
		os.Exit(0)
	}()

}
