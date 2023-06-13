package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log = logInit()

func logInit() *zap.Logger {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	level, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = zap.InfoLevel
	}
	return zap.New(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(colorable.NewColorableStdout()), level),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
}

func Info(msg string, field ...zap.Field) {
	log.Info(msg, field...)
}

func Debug(msg string, field ...zap.Field) {
	log.Debug(msg, field...)
}

func Error(msg string, field ...zap.Field) {
	log.Error(msg, field...)
}

func Fatal(msg string, field ...zap.Field) {
	log.Fatal(msg, field...)
}
