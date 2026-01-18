package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(environment string) (*Logger, error) {
	var config zap.Config

	if environment == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config.OutputPaths = []string{"stdout", "./storage/logs/app.log"}
	config.ErrorOutputPaths = []string{"stderr", "./storage/logs/error.log"}

	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}

	sugar := zapLogger.Sugar()
	return &Logger{sugar}, nil
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
