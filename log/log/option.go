package log

type Option func(g *option)

type option struct {
	syncFile    string
	syncConsole bool
	//syncEmail   string //同步邮箱
	//syncHttp    string //同步http
	timeFormat      string
	maxFileSize     int //日志文件最大多大
	maxAge          int //文件最多保存多少天
	formatJsonStyle bool
	depth           int
}

func WithSyncFile(file string) Option {
	return func(o *option) {
		o.syncFile = file
	}
}

func WithSyncConsole() Option {
	return func(o *option) {
		o.syncConsole = true
	}
}

func WithCallDepth(depth int) Option {
	return func(o *option) {
		o.depth = depth
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
