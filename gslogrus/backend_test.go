package gslogrus

import (
	"testing"
	"time"

	"github.com/hindsights/gslog"
	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		// DisableQuote:   true,
		DisableSorting: true,
	})
	gslog.SetBackend(NewBackend(logrusLogger, gslog.LogLevelAll))
	gslog.Info("gs-hello")
	gslog.Warn("start")
	logger := gslog.GetLogger("app")
	for {
		logger.Debug("debug", 1)
		logger.Info("info", "abc")
		logger.Warn("warn", true)
		logger.Error("error", false)
		logger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Error("field output")
		logger.Log(gslog.LogLevelDebug, "log debug", "value1", "value2")
		logger.Logf(gslog.LogLevelDebug, "log debug format key1=%s key2=%d", "value1", 123)
		time.Sleep(time.Second)
		gslog.Debugf("debugf %s", "name")
		gslog.Infof("infof %s", "value")
		gslog.Warnf("warnf %d", 20)
		gslog.Errorf("errorf %v", 100)
		time.Sleep(time.Second * 3)
		break
	}
}
