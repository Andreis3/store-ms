package logger_mock

type LoggerMock struct {
	InfoMSG   string
	DebugMSG  string
	WarnMSG   string
	ErrorMSG  string
	InfoAny   any
	DebugAny  any
	WarnAny   any
	ErrorAny  any
	InfoFunc  func(msg string, info ...any)
	DebugFunc func(msg string, info ...any)
	WarnFunc  func(msg string, info ...any)
	ErrorFunc func(msg string, info ...any)
}

func (l *LoggerMock) Info(msg string, info ...any) {
	l.InfoMSG = msg
	l.InfoAny = info
	l.InfoFunc(msg, info...)
}

func (l *LoggerMock) Debug(msg string, info ...any) {
	l.DebugMSG = msg
	l.DebugAny = info
	l.DebugFunc(msg, info...)
}

func (l *LoggerMock) Warn(msg string, info ...any) {
	l.WarnMSG = msg
	l.WarnAny = info
	l.WarnFunc(msg, info...)
}

func (l *LoggerMock) Error(msg string, info ...any) {
	l.ErrorMSG = msg
	l.ErrorAny = info
	l.ErrorFunc(msg, info...)
}
