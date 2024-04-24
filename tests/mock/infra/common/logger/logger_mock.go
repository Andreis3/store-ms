package logger_mock

import "github.com/stretchr/testify/mock"

const (
	Info  = "Info"
	Debug = "Debug"
	Warn  = "Warn"
	Error = "Error"
)

type LoggerMock struct {
	mock.Mock
}

func (l *LoggerMock) Info(msg string, info ...any) {
	l.Called(msg, info)
}

func (l *LoggerMock) Debug(msg string, info ...any) {
	l.Called(msg, info)
}

func (l *LoggerMock) Warn(msg string, info ...any) {
	l.Called(msg, info)
}

func (l *LoggerMock) Error(msg string, info ...any) {
	l.Called(msg, info)
}
