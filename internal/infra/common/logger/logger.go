package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	logger slog.Logger
}

func NewLogger() *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	slog.SetDefault(logger)
	return &Logger{*logger}
}

func (l *Logger) Debug(msg string, info any) {
	l.logger.Debug(msg, slog.String("info", fmt.Sprintf("%s", info)))
}
func (l *Logger) Info(msg string, info any) {

	l.logger.Info(msg, slog.String("info", fmt.Sprintf("%s", info)))
}
func (l *Logger) Warn(msg string, info any) {
	l.logger.Warn(msg, slog.String("info", fmt.Sprintf("%s", info)))
}
func (l *Logger) Error(msg string, info any) {
	l.logger.Error(msg, slog.String("info", fmt.Sprintf("%s", info)))
}
