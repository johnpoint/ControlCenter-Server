package log

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

// WithJSONEncoding 可以指定 json 编码方式
func WithJSONEncoding() Option {
	return optionFunc(func(log *Logger) {
		log.encoding = "json"
	})
}

// WithConsoleEncoding 可以指定 console 编码方式
func WithConsoleEncoding() Option {
	return optionFunc(func(log *Logger) {
		log.encoding = "console"
	})
}

func WrapLevelEncoder(levelEncoder LevelEncoder) Option {
	return optionFunc(func(logger *Logger) {
		logger.levelEncoder = levelEncoder
	})
}

func IsDev() Option {
	return optionFunc(func(logger *Logger) {
		logger.isDev = true
	})
}
