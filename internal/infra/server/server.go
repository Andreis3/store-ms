package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/common/configs"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/proxy"
	"github.com/andreis3/stores-ms/internal/util"
)

func Start(conf *configs.Conf, log *logger.Logger) {
	start := time.Now()
	mux := chi.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", conf.ServerPort),
		Handler: mux,
	}
	pool := postgres.NewPostgresDB(*conf)
	go func() {
		proxy.ProxyDependency(mux, pool, log)
		end := time.Now()
		microseconds := end.Sub(start).Microseconds()
		log.Info(fmt.Sprintf("Server started in %d microseconds", microseconds))
		log.Info(fmt.Sprintf("Start server on port %s", conf.ServerPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
			os.Exit(util.EXIT_FAILURE)
		}
	}()
	gracefulShutdown(server, pool, log)
}
