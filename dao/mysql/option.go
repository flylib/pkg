package mysql

import "regexp"

type Option func(*option)

// LogLevel log level
type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
)

type option struct {
	maxOpenConns int
	maxIdleConns int
	logLevel     LogLevel
}

func WithMaxOpenConns(num int) Option {
	return func(o *option) {
		o.maxOpenConns = num
	}
}

func WithLogLevel(lv LogLevel) Option {
	return func(o *option) {
		o.logLevel = lv
	}
}

// The default max idle connections is currently 2.
func WithMaxIdleConns(num int) Option {
	return func(o *option) {
		o.maxIdleConns = num
	}
}

func WithRawParamRegexp(exp string) Option {
	return func(o *option) {
		var err error
		rawRegexp, err = regexp.Compile(exp)
		if err != nil {
			panic(err)
		}
	}
}

func WithProcedureParamRegexp(exp string) Option {
	return func(o *option) {
		var err error
		pcdRegexp, err = regexp.Compile(exp)
		if err != nil {
			panic(err)
		}
	}
}
