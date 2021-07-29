package gslogrus

import (
	"fmt"

	"github.com/hindsights/gslog"
)

type logrusLogger struct {
	backend *logursBackend
	name    string
}

// func (logger *logrusLogger) Name() string {
// 	return logger.name
// }

func newLogger(backend *logursBackend, name string) gslog.Logger {
	return &logrusLogger{
		backend: backend,
		name:    name,
	}
}

func (logger *logrusLogger) formatLoggerName() string {
	return fmt.Sprintf("[%6s]", logger.name)
}

func (logger *logrusLogger) prepareArgs(args []interface{}) []interface{} {
	newArgs := make([]interface{}, len(args)*2+1)
	newArgs[0] = logger.formatLoggerName()
	for i, arg := range args {
		// add extra space separator
		newArgs[i*2+1] = " "
		newArgs[i*2+2] = arg
	}
	return newArgs
}

func (logger *logrusLogger) prepareFormatArgs(format string, args []interface{}) (string, []interface{}) {
	newArgs := make([]interface{}, len(args)+1)
	newArgs[0] = logger.formatLoggerName()
	for i, arg := range args {
		newArgs[i+1] = arg
	}
	return "%s " + format, newArgs
}

func (logger *logrusLogger) doLog(level gslog.LogLevel, f func(...interface{}), args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newArgs := logger.prepareArgs(args)
	f(newArgs...)
}

func (logger *logrusLogger) doLogf(level gslog.LogLevel, f func(string, ...interface{}), format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newFormat, newArgs := logger.prepareFormatArgs(format, args)
	f(newFormat, newArgs...)
}

func (logger *logrusLogger) NeedLog(level gslog.LogLevel) bool {
	return logger.backend.logger.IsLevelEnabled(FromGSLogLevel(level))
}

func (logger *logrusLogger) Logf(level gslog.LogLevel, format string, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newFormat, newArgs := logger.prepareFormatArgs(format, args)
	logger.backend.logger.Logf(FromGSLogLevel(level), newFormat, newArgs...)
}

func (logger *logrusLogger) Log(level gslog.LogLevel, args ...interface{}) {
	if !logger.NeedLog(level) {
		return
	}
	newArgs := logger.prepareArgs(args)
	logger.backend.logger.Log(FromGSLogLevel(level), newArgs...)
}

func (logger *logrusLogger) Trace(args ...interface{}) {
	logger.doLog(gslog.LogLevelTrace, logger.backend.logger.Trace, args...)
}

func (logger *logrusLogger) Debug(args ...interface{}) {
	logger.doLog(gslog.LogLevelDebug, logger.backend.logger.Debug, args...)
}

func (logger *logrusLogger) Info(args ...interface{}) {
	logger.doLog(gslog.LogLevelInfo, logger.backend.logger.Info, args...)
}

func (logger *logrusLogger) Warn(args ...interface{}) {
	logger.doLog(gslog.LogLevelWarn, logger.backend.logger.Warn, args...)
}

func (logger *logrusLogger) Error(args ...interface{}) {
	logger.doLog(gslog.LogLevelError, logger.backend.logger.Error, args...)
}

func (logger *logrusLogger) Fatal(args ...interface{}) {
	logger.doLog(gslog.LogLevelFatal, logger.backend.logger.Fatal, args...)
}

func (logger *logrusLogger) Tracef(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelTrace, logger.backend.logger.Tracef, format, args...)
}

func (logger *logrusLogger) Debugf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelDebug, logger.backend.logger.Debugf, format, args...)
}

func (logger *logrusLogger) Infof(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelInfo, logger.backend.logger.Infof, format, args...)
}

func (logger *logrusLogger) Warnf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelWarn, logger.backend.logger.Warnf, format, args...)
}

func (logger *logrusLogger) Errorf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelError, logger.backend.logger.Errorf, format, args...)
}

func (logger *logrusLogger) Fatalf(format string, args ...interface{}) {
	logger.doLogf(gslog.LogLevelFatal, logger.backend.logger.Fatalf, format, args...)
}

func (logger *logrusLogger) WithFields(fields gslog.Fields) gslog.Logger {
	return newLoggerEntry(logger, fields)
}
