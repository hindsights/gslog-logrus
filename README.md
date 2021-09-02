# gslog-logrus

gslog backend based on logrus library

## Example

```go
package main

import (
	"fmt"

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
	gslog.SetBackend(NewBackend(logrusLogger))
	gslog.Info("gs-hello")
	gslog.Warn("start")
	logger := gslog.GetLogger("app")
	flogger := gslog.GetFieldLogger("app")

	flogger.Debug("debug", gslog.Fields{"integer": 1})
	flogger.Info("info", gslog.Fields{"abc": 234})
	flogger.Warn("warn", gslog.Fields{"bool": true})
	flogger.Error("error", gslog.Fields{"bool": false})
	flogger.Log(gslog.LogLevelDebug, "log debug", gslog.Fields{"value1": "value2"})
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

```
