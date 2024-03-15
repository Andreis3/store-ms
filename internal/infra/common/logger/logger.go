package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger slog.Logger
}

func NewLogger() *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	slog.SetDefault(logger)
	return &Logger{*logger}
}

func (l *Logger) Debug(msg string, info map[string]any) {
	l.logger.Debug(msg, info)
}
func (l *Logger) Info(msg string, info map[string]any) {
	l.logger.Info(msg, info)
}
func (l *Logger) Warn(msg string, info map[string]any) {
	l.logger.Warn(msg, info)
}
func (l *Logger) Error(msg string, info map[string]any) {
	l.logger.Error(msg, info)
}
