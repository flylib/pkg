package zaplog

import (
	"testing"
)

func TestZaplog(t *testing.T) {
	logger := NewZapLogger(
		WithSyncFile("./info.log"),
		WithSyncConsole(),
		WithCallDepth(1),
		//MinPrintLevel(DebugLevel),
	)

	for i := 0; i < 1; i++ {
		logger.Debug("hello", i)
		logger.Warn("hello", i)
		logger.Infof("hello %d", i)
		logger.Warnf("hello %d", i)
		logger.Errorf("hello %d", i)
	}
}
