package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field
type Level = zapcore.Level

var l = NewDefaultLogger()

var (
	Info  = l.Info
	Error = l.Error
	Warn  = l.Warn
	Panic = l.Panic
)

type Config = zap.Config
type LevelEncoder = zapcore.LevelEncoder
type EncoderConfig = zapcore.EncoderConfig

var (
	String = zap.String
)

type Logger struct {
	l            *zap.Logger
	level        Level
	encoding     string
	levelEncoder LevelEncoder
	isDev        bool
}

func OverrideLoggerWithOption(keyValue map[string]interface{}, options ...Option) {
	for i := range options {
		options[i].apply(l)
	}
	loggerConfig := Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochNanosTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    keyValue,
	}

	if len(l.encoding) != 0 {
		loggerConfig.Encoding = l.encoding
	}
	if l.levelEncoder != nil {
		loggerConfig.EncoderConfig.EncodeLevel = l.levelEncoder
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	l.l = logger
}

func NewDefaultLogger() *Logger {
	logger, err := Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochNanosTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		panic(err)
	}
	return &Logger{
		l: logger,
	}
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.level = zap.InfoLevel
	l.l.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.level = zap.ErrorLevel
	l.l.Error(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.level = zap.WarnLevel
	l.l.Warn(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.level = zap.PanicLevel
	l.l.Panic(msg, fields...)
}
