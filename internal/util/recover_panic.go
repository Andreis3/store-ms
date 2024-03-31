package util

import (
	"fmt"
	"os"

	"github.com/andreis3/stores-ms/internal/infra/common/logger"
)

func RecoverFromPanic(log *logger.Logger) {
	if r := recover(); r != nil {
		log.Error(fmt.Sprintf("Recovered from panic: %v", r))
		os.Exit(EXIT_FAILURE)
	}
}
