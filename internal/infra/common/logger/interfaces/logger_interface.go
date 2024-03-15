package ilogger

type ILogger interface {
	Debug(msg string, info map[string]any)
	Info(msg string, info map[string]any)
	Warn(msg string, info map[string]any)
	Error(msg string, info map[string]any)
}
