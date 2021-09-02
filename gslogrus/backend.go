package gslogrus

import (
	"github.com/hindsights/gslog"
	"github.com/sirupsen/logrus"
)

func FromGSLogLevel(level gslog.LogLevel) logrus.Level {
	switch {
	case level == gslog.LogLevelDebug:
		return logrus.DebugLevel
	case level == gslog.LogLevelInfo:
		return logrus.InfoLevel
	case level == gslog.LogLevelWarn:
		return logrus.WarnLevel
	case level == gslog.LogLevelError:
		return logrus.ErrorLevel
	case level >= gslog.LogLevelFatal:
		return logrus.FatalLevel
	}
	return logrus.FatalLevel
}

func ToGSLogLevel(level logrus.Level) gslog.LogLevel {
	switch {
	case level == logrus.DebugLevel:
		return gslog.LogLevelDebug
	case level == logrus.InfoLevel:
		return gslog.LogLevelInfo
	case level == logrus.WarnLevel:
		return gslog.LogLevelWarn
	case level == logrus.ErrorLevel:
		return gslog.LogLevelError
	case level >= logrus.FatalLevel:
		return gslog.LogLevelFatal
	}
	return gslog.LogLevelFatal
}

type logursBackend struct {
	logger *logrus.Logger
}

func (backend *logursBackend) GetLogger(name string) gslog.Logger {
	return newLogger(backend, name)
}

func (backend *logursBackend) GetFieldLogger(name string) gslog.FieldLogger {
	return newSLogger(backend, name)
}

func NewBackend(logger *logrus.Logger) gslog.Backend {
	return &logursBackend{logger: logger}
}
