package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var err error
	if logger, err = config.Build(); err != nil {
		panic(err)
	}
}

func Error(message string, err error) {
	logger.Error(message, zap.Error(err))
}

func Info(message string) {
	logger.Info(message)
}

func Debug(message string) {
	logger.Debug(message)
}

func Warn(message string) {
	logger.Warn(message)
}

func Fatal(message string, err error) {
	logger.Fatal(message, zap.Error(err))
}
