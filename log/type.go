package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	Level zap.AtomicLevel
}

func (L *Logger) SetLogLevel(level zapcore.Level) {
	L.Level.SetLevel(level)
}
