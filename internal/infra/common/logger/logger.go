package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type Logger struct {
	logger slog.Logger
}

func NewLogger() *Logger {
	w := os.Stderr
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.DateTime,
			NoColor:    false,
		}),
	)
	slog.SetDefault(logger)
	return &Logger{*logger}
}
func (l *Logger) Debug(msg string, info ...any) {
	l.logger.Debug(msg, info...)
}
func (l *Logger) Info(msg string, info ...any) {

	l.logger.Info(msg, info...)
}
func (l *Logger) Warn(msg string, info ...any) {
	l.logger.Warn(msg, info...)
}
func (l *Logger) Error(msg string, info ...any) {
	l.logger.Error(msg, info...)
}
