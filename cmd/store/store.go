package main

import (
	"fmt"
	"os"

	"github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres"
	"github.com/andreis3/stores-ms/internal/infra/common/configs"
	"github.com/andreis3/stores-ms/internal/infra/common/logger"
	"github.com/andreis3/stores-ms/internal/infra/server"
	"github.com/andreis3/stores-ms/internal/util"
)

func main() {
	log := logger.NewLogger()
	conf, err := configs.LoadConfig(".")
	if err != nil {
		log.Error(fmt.Sprintf("Error loading config: %s", err.Error()))
		os.Exit(util.EXIT_FAILURE)
	}
	pool := postgres.NewPostgresDB(*conf)
	server.GracefulShutdown(server.Start(conf, pool, log), pool, log)
}
