package zaplog

import (
	"testing"
)

func TestZaplog(t *testing.T) {
	sugaredLogger := NewZapLogger(
		WithSyncFile("./info.log"),
		WithSyncConsole(),
		//MinPrintLevel(DebugLevel),
	).Sugar()

	for i := 0; i < 10; i++ {
		sugaredLogger.Info("hello", i)
		sugaredLogger.Debug("hello", i)
		sugaredLogger.Infof("hello %d", i)
	}
}
