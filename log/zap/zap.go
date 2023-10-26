package zaplog

import (
	Ilog "github.com/flylib/interface/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"os"
	"time"
)

func NewZapLogger(options ...Option) *zap.Logger {
	opt := option{}
	for _, f := range options {
		f(&opt)
	}

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		if opt.timeFormat == "" {
			opt.timeFormat = time.DateTime
		}
		enc.AppendString(t.Format(opt.timeFormat))
	}

	//log encoder config
	cfg := zapcore.EncoderConfig{
		CallerKey:      "caller", // 打印文件名和行数 json格式时生效
		LevelKey:       "lv",
		MessageKey:     "msg",
		TimeKey:        "time",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,           // 自定义时间格式
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder,  // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	if opt.formatJsonStyle {
		encoder = zapcore.NewJSONEncoder(cfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(cfg)
	}

	//日志打印级别
	level := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		if opt.minLogLevel >= Ilog.FatalLevel {
			opt.minLogLevel += 2
		}
		return lv >= zapcore.Level(opt.minLogLevel)
	})

	//同步输出
	var cores []zapcore.Core
	if opt.syncConsole || opt.syncFile == "" {
		cores = append(cores, zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stdout),
			level))
	}
	if opt.syncFile != "" {
		syncFile := zapcore.AddSync(&lumberjack.Logger{
			Filename:  opt.syncFile,
			MaxSize:   opt.maxFileSize,
			MaxAge:    opt.maxAge,
			LocalTime: true,
			Compress:  false,
		})
		cores = append(cores, zapcore.NewCore(
			encoder,
			syncFile,
			level))
	}

	return zap.New(zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(opt.depth))
}
