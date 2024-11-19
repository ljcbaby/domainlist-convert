package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger      *zap.Logger
	atomicLevel zap.AtomicLevel
)

func Init() {
	atomicLevel = zap.NewAtomicLevel()

	config := zap.Config{
		Level:       atomicLevel,
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
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	Logger, err = config.Build()
	SetLogLevel(zapcore.InfoLevel)
	if err != nil {
		panic(err)
	}
}

func SetLogLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}
