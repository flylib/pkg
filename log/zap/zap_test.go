package zaplog

import (
	Ilog "github.com/flylib/interface"
	"testing"
)

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(
		WithSyncFile("./info.log"),
		WithSyncConsole(),
		//MinPrintLevel(DebugLevel),
		WithMinLogLevel(Ilog.DebugLevel),
		WithColorEncoder(),
	)
	i := 1
	logger.Debug("debug", i)
	logger.Info("info", i)
	logger.Warn("warn", i)
	logger.Error("error", i)

	logger.Debugf("debug %d", i)
	logger.Infof("info %d", i)
	logger.Warnf("warn %d", i)
	logger.Errorf("error %d", i)
}
