package ilogger

type ILogger interface {
	Debug(msg string, info ...any)
	Info(msg string, info ...any)
	Warn(msg string, info ...any)
	Error(msg string, info ...any)
}
