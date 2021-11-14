package main

import (
	"fmt"
	"time"

	"github.com/hindsights/gslog"
	"github.com/hindsights/gslog-logrus/gslogrus"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("test")

	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors:  true,
		FullTimestamp:  true,
		DisableSorting: true,
	})
	logrusLogger.SetLevel(logrus.DebugLevel)
	gslog.SetBackend(gslogrus.NewBackend(logrusLogger))
	gslog.Info("gs-hello")
	gslog.Warn("start")
	logger := gslog.GetSimpleLogger("app")
	flogger := gslog.GetLogger("app")

	flogger.Int("int", 1).Debug("debug")
	flogger.Int("abc", 234).Info("info")
	flogger.Bool("bool", true).Warn("warn")
	flogger.Bool("bool", false).Error("error")
	flogger.Str("val", "val2").Log(gslog.LogLevelDebug, "log debug")
	time.Sleep(time.Second)
	logger.Debug("debug", "name")
	logger.Info("info", "value")
	logger.Warn("warn", 20)
	logger.Error("error", 100)
	logger.Debugf("debugf %s", "name")
	logger.Infof("infof %s", "value")
	logger.Warnf("warnf %d", 20)
	logger.Errorf("errorf %v", 100)
	time.Sleep(time.Second * 3)
}
