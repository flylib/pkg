package log

import "testing"

func TestLogger(t *testing.T) {
	//newLogger := NewLogger(SyncFile("./log.log"), MaxFileSize(1))
	newLogger := NewLogger(WithSyncConsole())

	for i := 0; i < 10; i++ {
		newLogger.Info("info")
		newLogger.Infof("info %d", 123)
		newLogger.Errorf("test %d", 123)
	}

}
