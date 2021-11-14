package gslogrus

import (
	"time"

	"github.com/hindsights/gslog"
	"github.com/sirupsen/logrus"
)

type fieldLogger struct {
	backend *logursBackend
	name    string
	entry   *logrus.Entry
}

func newFieldLogger(backend *logursBackend, name string) gslog.Logger {
	return fieldLogger{
		backend: backend,
		name:    name,
		entry:   backend.logger.WithField("ctx", name),
	}
}

func (logger fieldLogger) NeedLog(level gslog.LogLevel) bool {
	return logger.backend.logger.IsLevelEnabled(FromGSLogLevel(level))
}

func (logger fieldLogger) Fields(fields gslog.Fields) gslog.Logger {
	return fieldLogger{backend: logger.backend, name: logger.name, entry: logger.entry.WithFields(logrus.Fields(fields))}
}

func (logger fieldLogger) Field(key string, val interface{}) gslog.Logger {
	return fieldLogger{backend: logger.backend, name: logger.name, entry: logger.entry.WithFields(logrus.Fields{key: val})}
}

func (logger fieldLogger) Log(level gslog.LogLevel, msg string) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Log(FromGSLogLevel(level), msg)
}

func (logger fieldLogger) Debug(msg string) {
	logger.Log(gslog.LogLevelDebug, msg)
}

func (logger fieldLogger) Info(msg string) {
	logger.Log(gslog.LogLevelInfo, msg)
}

func (logger fieldLogger) Warn(msg string) {
	logger.Log(gslog.LogLevelWarn, msg)
}

func (logger fieldLogger) Error(msg string) {
	logger.Log(gslog.LogLevelError, msg)
}

func (logger fieldLogger) Fatal(msg string) {
	logger.Log(gslog.LogLevelFatal, msg)
}

func (logger fieldLogger) Str(key string, val string) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Int(key string, val int) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Uint(key string, val uint) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Bool(key string, val bool) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Int64(key string, val int64) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Int32(key string, val int32) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Int16(key string, val int16) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Int8(key string, val int8) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Uint64(key string, val uint64) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Uint32(key string, val uint32) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Uint16(key string, val uint16) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Uint8(key string, val uint8) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Float32(key string, val float32) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Float64(key string, val float64) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Err(key string, val error) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Time(key string, val time.Time) gslog.Logger {
	return logger.Field(key, val)
}

func (logger fieldLogger) Duration(key string, val time.Duration) gslog.Logger {
	return logger.Field(key, val)
}

type sugaredLogger struct {
	backend *logursBackend
	name    string
	entry   *logrus.Entry
}

func newSugaredLogger(backend *logursBackend, name string) gslog.SugaredLogger {
	return sugaredLogger{
		backend: backend,
		name:    name,
		entry:   backend.logger.WithField("ctx", name),
	}
}

func (logger sugaredLogger) prepareArgs(args []interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}
	newArgs := make([]interface{}, len(args)*2-1)
	for i, arg := range args {
		// add extra space separator
		newArgs[i*2] = arg
		if i+1 < len(args) {
			newArgs[i*2+1] = " "
		}
	}
	return newArgs
}

func (logger sugaredLogger) doLog(level gslog.LogLevel, f func(...interface{}), args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newArgs := logger.prepareArgs(args)
	f(newArgs...)
}

func (logger sugaredLogger) doLogf(level gslog.LogLevel, f func(string, ...interface{}), format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	f(format, args...)
}

func (logger sugaredLogger) NeedLog(level gslog.LogLevel) bool {
	return logger.backend.logger.IsLevelEnabled(FromGSLogLevel(level))
}

func (logger sugaredLogger) Logf(level gslog.LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Logf(FromGSLogLevel(level), format, args...)
}

func (logger sugaredLogger) Log(level gslog.LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Log(FromGSLogLevel(level), args...)
}

func (logger sugaredLogger) Debug(args ...interface{}) {
	logger.doLog(gslog.LogLevelDebug, logger.entry.Debug, args...)
}

func (logger sugaredLogger) Info(args ...interface{}) {
	logger.doLog(gslog.LogLevelInfo, logger.entry.Info, args...)
}

func (logger sugaredLogger) Warn(args ...interface{}) {
	logger.doLog(gslog.LogLevelWarn, logger.entry.Warn, args...)
}

func (logger sugaredLogger) Error(args ...interface{}) {
	logger.doLog(gslog.LogLevelError, logger.entry.Error, args...)
}

func (logger sugaredLogger) Fatal(args ...interface{}) {
	logger.doLog(gslog.LogLevelFatal, logger.entry.Fatal, args...)
}

func (logger sugaredLogger) Debugf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelDebug, logger.entry.Debugf, format, args...)
}

func (logger sugaredLogger) Infof(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelInfo, logger.entry.Infof, format, args...)
}

func (logger sugaredLogger) Warnf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelWarn, logger.entry.Warnf, format, args...)
}

func (logger sugaredLogger) Errorf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelError, logger.entry.Errorf, format, args...)
}

func (logger sugaredLogger) Fatalf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelFatal, logger.entry.Fatalf, format, args...)
}
