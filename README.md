# gslog-logrus

gslog backedn based on logrus

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
		DisableColors: true,
		FullTimestamp: true,
		// DisableQuote:   true,
		DisableSorting: true,
	})
	gslog.SetBackend(gslogrus.NewBackend(logrusLogger, gslog.LogLevelAll))
	gslog.Info("gs-logrus-hello")
	gslog.Info("start")
	logger := gslog.GetLogger("app")
	logger.Debug("debug", 1)
	logger.Info("info", "abc")
	logger.Warn("warn", true)
	logger.Error("error", false)
	logger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Error("field output")
	logger.WithFields(gslog.Fields{"key1": 1, "key2": "val2"}).Errorf("field output %d", 567)
	gslog.Debugf("debugf %s", "name")
	gslog.Infof("infof %s", "value")
	gslog.Warnf("warnf %d", 20)
	gslog.Errorf("errorf %v", 100)
	logger.Info("output to logrus")
}

```