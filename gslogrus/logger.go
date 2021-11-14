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

type simpleLogger struct {
	backend *logursBackend
	name    string
	entry   *logrus.Entry
}

func newSimpleLogger(backend *logursBackend, name string) gslog.SimpleLogger {
	return simpleLogger{
		backend: backend,
		name:    name,
		entry:   backend.logger.WithField("ctx", name),
	}
}

func (logger simpleLogger) prepareArgs(args []interface{}) []interface{} {
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

func (logger simpleLogger) doLog(level gslog.LogLevel, f func(...interface{}), args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newArgs := logger.prepareArgs(args)
	f(newArgs...)
}

func (logger simpleLogger) doLogf(level gslog.LogLevel, f func(string, ...interface{}), format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	f(format, args...)
}

func (logger simpleLogger) NeedLog(level gslog.LogLevel) bool {
	return logger.backend.logger.IsLevelEnabled(FromGSLogLevel(level))
}

func (logger simpleLogger) Logf(level gslog.LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Logf(FromGSLogLevel(level), format, args...)
}

func (logger simpleLogger) Log(level gslog.LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Log(FromGSLogLevel(level), args...)
}

func (logger simpleLogger) Debug(args ...interface{}) {
	logger.doLog(gslog.LogLevelDebug, logger.entry.Debug, args...)
}

func (logger simpleLogger) Info(args ...interface{}) {
	logger.doLog(gslog.LogLevelInfo, logger.entry.Info, args...)
}

func (logger simpleLogger) Warn(args ...interface{}) {
	logger.doLog(gslog.LogLevelWarn, logger.entry.Warn, args...)
}

func (logger simpleLogger) Error(args ...interface{}) {
	logger.doLog(gslog.LogLevelError, logger.entry.Error, args...)
}

func (logger simpleLogger) Fatal(args ...interface{}) {
	logger.doLog(gslog.LogLevelFatal, logger.entry.Fatal, args...)
}

func (logger simpleLogger) Debugf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelDebug, logger.entry.Debugf, format, args...)
}

func (logger simpleLogger) Infof(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelInfo, logger.entry.Infof, format, args...)
}

func (logger simpleLogger) Warnf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelWarn, logger.entry.Warnf, format, args...)
}

func (logger simpleLogger) Errorf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelError, logger.entry.Errorf, format, args...)
}

func (logger simpleLogger) Fatalf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelFatal, logger.entry.Fatalf, format, args...)
}
