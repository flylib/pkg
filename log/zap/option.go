package zaplog

import Ilog "github.com/flylib/interface"

type Option func(o *option)

// options
type option struct {
	syncFile    string
	syncConsole bool
	//syncEmail   string //同步邮箱
	//syncHttp    string //同步http
	timeFormat      string
	maxFileSize     int //日志文件最大多大
	maxAge          int //文件最多保存多少天
	formatJsonStyle bool
	minLogLevel     Ilog.Level
	depth           int //default 1
}

// 同步写入文件
func WithSyncFile(file string) Option {
	return func(o *option) {
		o.syncFile = file
	}
}

// 是否同步控制台
func WithSyncConsole() Option {
	return func(o *option) {
		o.syncConsole = true
	}
}

// 时间格式
func WithTimeFormat(format string) Option {
	return func(o *option) {
		o.timeFormat = format
	}
}

// 单个日志文件大小（单位:MB）
func WithMaxFileSize(size int) Option {
	return func(o *option) {
		o.maxFileSize = size
	}
}

// 文件最多保留多长时间(单位:Day)
func WithMaxSaveDuration(day int) Option {
	return func(o *option) {
		o.maxAge = day
	}
}

// 输出jason格式
func WithOutJsonCodec() Option {
	return func(o *option) {
		o.formatJsonStyle = true
	}
}

// 最低打印日志级别
func WithMinLogLevel(lv Ilog.Level) Option {
	return func(o *option) {
		o.minLogLevel = lv - 1
	}
}

func WithCallDepth(depth int) Option {
	return func(o *option) {
		o.depth = depth
	}
}
