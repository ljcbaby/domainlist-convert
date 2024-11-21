package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *Logger = nil

func L() *Logger {
	if l == nil {
		l, _ = Init()
	}
	return l
}

func Init() (*Logger, error) {
	l := Logger{}

	l.Level = zap.NewAtomicLevel()

	config := zap.Config{
		Level:       l.Level,
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
	}

	var err error
	l.Logger, err = config.Build()
	l.SetLogLevel(zapcore.InfoLevel)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
