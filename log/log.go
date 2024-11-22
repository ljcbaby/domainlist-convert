package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	l  *Logger = Init()
	mu sync.RWMutex
)

func L() *Logger {
	mu.RLock()
	t := l
	mu.RUnlock()
	return t
}

func Init() *Logger {
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
	if err != nil {
		panic(err)
	}
	l.SetLogLevel(zapcore.InfoLevel)

	return &l
}
