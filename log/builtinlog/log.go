package builtinlog

import (
	"fmt"
	Ilog "github.com/flylib/interface/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	callDepth int
}

func NewLogger(options ...Option) *Logger {
	newLogger := new(Logger)
	opt := option{depth: 2}
	for _, f := range options {
		f(&opt)
	}
	newLogger.callDepth = opt.depth

	//print to console
	if opt.syncConsole || opt.syncFile == "" {
		newLogger.Logger = log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)
	} else {
		newLogger.Logger = log.New(&lumberjack.Logger{
			Filename:  opt.syncFile,
			MaxSize:   opt.maxFileSize,
			MaxAge:    opt.maxAge,
			LocalTime: true,
			Compress:  false,
		}, "", log.Lshortfile|log.LstdFlags)
	}
	return newLogger
}

func (l *Logger) Debug(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.DebugLevel), fmt.Sprint(args...)))
}

func (l *Logger) Info(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.InfoLevel), fmt.Sprint(args...)))
}

func (l *Logger) Warn(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.WarnLevel), fmt.Sprint(args...)))
}

func (l *Logger) Error(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.ErrorLevel), fmt.Sprint(args...)))
}

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.FatalLevel), fmt.Sprint(args...)))
	os.Exit(1)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.DebugLevel), fmt.Sprintf(format, args...)))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.InfoLevel), fmt.Sprintf(format, args...)))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.WarnLevel), fmt.Sprintf(format, args...)))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.ErrorLevel), fmt.Sprintf(format, args...)))
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, args ...any) {
	l.Output(l.callDepth, fmt.Sprintf("[%s] %s", Ilog.LevelString(Ilog.FatalLevel), fmt.Sprintf(format, args...)))
}
