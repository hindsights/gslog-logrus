package gslogrus

import (
	"github.com/hindsights/gslog"
	"github.com/sirupsen/logrus"
)

type sLogger struct {
	backend *logursBackend
	name    string
	entry   *logrus.Entry
}

func newSLogger(backend *logursBackend, name string) gslog.FieldLogger {
	return sLogger{
		backend: backend,
		name:    name,
		entry:   backend.logger.WithField("ctx", name),
	}
}

func (logger sLogger) NeedLog(level gslog.LogLevel) bool {
	return logger.backend.logger.IsLevelEnabled(FromGSLogLevel(level))
}

func (logger sLogger) WithFields(fields gslog.Fields) gslog.FieldLogger {
	return sLogger{backend: logger.backend, name: logger.name, entry: logger.entry.WithFields(logrus.Fields(fields))}
}

func (logger sLogger) Log(level gslog.LogLevel, msg string, fields ...gslog.Fields) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.WithFields(logrus.Fields(gslog.JoinFields(fields...))).Log(FromGSLogLevel(level), msg)
}

func (logger sLogger) Debug(msg string, fields ...gslog.Fields) {
	logger.Log(gslog.LogLevelDebug, msg, fields...)
}

func (logger sLogger) Info(msg string, fields ...gslog.Fields) {
	logger.Log(gslog.LogLevelInfo, msg, fields...)
}

func (logger sLogger) Warn(msg string, fields ...gslog.Fields) {
	logger.Log(gslog.LogLevelWarn, msg, fields...)
}

func (logger sLogger) Error(msg string, fields ...gslog.Fields) {
	logger.Log(gslog.LogLevelError, msg, fields...)
}

func (logger sLogger) Fatal(msg string, fields ...gslog.Fields) {
	logger.Log(gslog.LogLevelFatal, msg, fields...)
}

type logrusLogger struct {
	backend *logursBackend
	name    string
	entry   *logrus.Entry
}

func newLogger(backend *logursBackend, name string) gslog.Logger {
	return logrusLogger{
		backend: backend,
		name:    name,
		entry:   backend.logger.WithField("ctx", name),
	}
}

func (logger logrusLogger) prepareArgs(args []interface{}) []interface{} {
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

func (logger logrusLogger) doLog(level gslog.LogLevel, f func(...interface{}), args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newArgs := logger.prepareArgs(args)
	f(newArgs...)
}

func (logger logrusLogger) doLogf(level gslog.LogLevel, f func(string, ...interface{}), format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	f(format, args...)
}

func (logger logrusLogger) NeedLog(level gslog.LogLevel) bool {
	return logger.backend.logger.IsLevelEnabled(FromGSLogLevel(level))
}

func (logger logrusLogger) Logf(level gslog.LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Logf(FromGSLogLevel(level), format, args...)
}

func (logger logrusLogger) Log(level gslog.LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	logger.entry.Log(FromGSLogLevel(level), args...)
}

func (logger logrusLogger) Debug(args ...interface{}) {
	logger.doLog(gslog.LogLevelDebug, logger.entry.Debug, args...)
}

func (logger logrusLogger) Info(args ...interface{}) {
	logger.doLog(gslog.LogLevelInfo, logger.entry.Info, args...)
}

func (logger logrusLogger) Warn(args ...interface{}) {
	logger.doLog(gslog.LogLevelWarn, logger.entry.Warn, args...)
}

func (logger logrusLogger) Error(args ...interface{}) {
	logger.doLog(gslog.LogLevelError, logger.entry.Error, args...)
}

func (logger logrusLogger) Fatal(args ...interface{}) {
	logger.doLog(gslog.LogLevelFatal, logger.entry.Fatal, args...)
}

func (logger logrusLogger) Debugf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelDebug, logger.entry.Debugf, format, args...)
}

func (logger logrusLogger) Infof(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelInfo, logger.entry.Infof, format, args...)
}

func (logger logrusLogger) Warnf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelWarn, logger.entry.Warnf, format, args...)
}

func (logger logrusLogger) Errorf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelError, logger.entry.Errorf, format, args...)
}

func (logger logrusLogger) Fatalf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelFatal, logger.entry.Fatalf, format, args...)
}
