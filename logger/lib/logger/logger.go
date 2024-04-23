package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var loggerClient *zap.Logger

func InitializeLogger() (*zap.Logger, error) {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Getting and using a value from .env
	log_level := os.Getenv("LOG_LEVEL")

	loggerConfig := zap.NewProductionConfig()
	level, _ := zapcore.ParseLevel(log_level)
	loggerConfig.Level = zap.NewAtomicLevelAt(level)
	loggerConfig.DisableCaller = true
	loggerConfig.DisableStacktrace = false
	loggerConfig.Sampling = nil
	loggerConfig.Encoding = "json"

	loggerConfig.InitialFields = getFields()

	loggerConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "messages",
		LevelKey:       "leval",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		return nil, err
	}
	logger.Debug("Logger initialized with sucess")
	return logger, nil

}

func getFields() map[string]interface{} {
	initialFields := map[string]interface{}{
		"application": "logger",
	}
	return initialFields
}

func GetClientLogger() *zap.Logger {
	if loggerClient == nil {
		loggerClient, _ = InitializeLogger()
	}

	return loggerClient
}

func Info(message string, fields ...interface{}) {
	logger := GetClientLogger()
	formattedMessage := fmt.Sprintf(message, fields...)
	logger.Info(formattedMessage)
}

func Error(message string, err error, fields ...interface{}) {
	logger := GetClientLogger()
	formattedMessage := fmt.Sprintf(message, fields...)
	if err != nil {
		logger.Error(formattedMessage)
	} else {
		logger.Error(formattedMessage, zap.Error(err))
	}

}

func Debug(message string, fields ...interface{}) {
	logger := GetClientLogger()
	formattedMessage := fmt.Sprintf(message, fields...)
	logger.Debug(formattedMessage)
}

func Warn(message string, fields ...interface{}) {
	logger := GetClientLogger()
	formattedMessage := fmt.Sprintf(message, fields...)
	logger.Warn(formattedMessage)
}

func Fatal(message string, fields ...interface{}) {
	logger := GetClientLogger()
	formattedMessage := fmt.Sprintf(message, fields...)
	logger.Fatal(formattedMessage)
}

func Panic(message string, fields ...interface{}) {
	logger := GetClientLogger()
	formattedMessage := fmt.Sprintf(message, fields...)
	logger.Panic(formattedMessage)
}
