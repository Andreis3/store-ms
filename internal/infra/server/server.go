package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/common/configs"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/proxy"
	"github.com/andreis3/stores-ms/internal/util"
)

func Start(conf *configs.Conf, pool *postgres.Postgres, log *logger.Logger) *http.Server {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", conf.ServerPort),
		Handler: mux,
	}
	go func() {
		proxy.ProxyDependency(mux, pool, log, conf)
		log.Info(fmt.Sprintf("Start server on port %s", conf.ServerPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
			os.Exit(util.EXIT_FAILURE)
		}
	}()
	return server
}
