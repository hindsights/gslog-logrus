package gslogrus

import (
	"github.com/hindsights/gslog"
	"github.com/sirupsen/logrus"
)

type logrusLoggerEntry struct {
	logger *logrusLogger
	entry  *logrus.Entry
}

func newLoggerEntry(logger *logrusLogger, fields gslog.Fields) gslog.Logger {
	return &logrusLoggerEntry{
		logger: logger,
		entry:  logger.backend.logger.WithFields(logrus.Fields(fields)),
	}
}

func (entry *logrusLoggerEntry) doLog(level gslog.LogLevel, f func(...interface{}), args ...interface{}) {
	if !entry.NeedLog(level) {
		return
	}
	newArgs := entry.logger.prepareArgs(args)
	f(newArgs...)
}

func (entry *logrusLoggerEntry) doLogf(level gslog.LogLevel, f func(string, ...interface{}), format string, args ...interface{}) {
	if !entry.NeedLog(level) {
		return
	}
	newFormat, newArgs := entry.logger.prepareFormatArgs(format, args)
	f(newFormat, newArgs...)
}

func (entry *logrusLoggerEntry) NeedLog(level gslog.LogLevel) bool {
	return entry.logger.NeedLog(level)
}

func (entry *logrusLoggerEntry) Logf(level gslog.LogLevel, format string, args ...interface{}) {
	if !entry.NeedLog(level) {
		return
	}
	newFormat, newArgs := entry.logger.prepareFormatArgs(format, args)
	entry.entry.Logf(FromGSLogLevel(level), newFormat, newArgs...)
}

func (entry *logrusLoggerEntry) Log(level gslog.LogLevel, args ...interface{}) {
	if !entry.logger.NeedLog(level) {
		return
	}
	newArgs := entry.logger.prepareArgs(args)
	entry.entry.Log(FromGSLogLevel(level), newArgs...)
}

func (entry *logrusLoggerEntry) Trace(args ...interface{}) {
	entry.doLog(gslog.LogLevelTrace, entry.entry.Trace, args...)
}

func (entry *logrusLoggerEntry) Debug(args ...interface{}) {
	entry.doLog(gslog.LogLevelDebug, entry.entry.Debug, args...)
}

func (entry *logrusLoggerEntry) Info(args ...interface{}) {
	entry.doLog(gslog.LogLevelInfo, entry.entry.Info, args...)
}

func (entry *logrusLoggerEntry) Warn(args ...interface{}) {
	entry.doLog(gslog.LogLevelWarn, entry.entry.Warn, args...)
}

func (entry *logrusLoggerEntry) Error(args ...interface{}) {
	entry.doLog(gslog.LogLevelError, entry.entry.Error, args...)
}

func (entry *logrusLoggerEntry) Fatal(args ...interface{}) {
	entry.doLog(gslog.LogLevelFatal, entry.entry.Fatal, args...)
}

func (entry *logrusLoggerEntry) Tracef(format string, args ...interface{}) {
	entry.doLogf(gslog.LogLevelTrace, entry.entry.Tracef, format, args...)
}

func (entry *logrusLoggerEntry) Debugf(format string, args ...interface{}) {
	entry.doLogf(gslog.LogLevelDebug, entry.entry.Debugf, format, args...)
}

func (entry *logrusLoggerEntry) Infof(format string, args ...interface{}) {
	entry.doLogf(gslog.LogLevelInfo, entry.entry.Infof, format, args...)
}

func (entry *logrusLoggerEntry) Warnf(format string, args ...interface{}) {
	entry.doLogf(gslog.LogLevelWarn, entry.entry.Warnf, format, args...)
}

func (entry *logrusLoggerEntry) Errorf(format string, args ...interface{}) {
	entry.doLogf(gslog.LogLevelError, entry.entry.Errorf, format, args...)
}

func (entry *logrusLoggerEntry) Fatalf(format string, args ...interface{}) {
	entry.doLogf(gslog.LogLevelFatal, entry.entry.Fatalf, format, args...)
}

func (entry *logrusLoggerEntry) WithFields(fields gslog.Fields) gslog.Logger {
	return newLoggerEntry(entry.logger, fields)
}
