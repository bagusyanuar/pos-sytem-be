package logger

import (
	"os"
	"path/filepath"

	"github.com/bagusyanuar/pos-sytem-be/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(cfg *config.LoggerConfig, environment string) (*Logger, error) {
	// Create logs directory from config path
	logDir := filepath.Dir(cfg.OutputPath)
	if err := ensureDir(logDir); err != nil {
		return nil, err
	}

	var zapConfig zap.Config

	if environment == "production" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.OutputPaths = []string{"stdout", cfg.OutputPath}
	zapConfig.ErrorOutputPaths = []string{"stderr", cfg.ErrorPath}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	sugar := zapLogger.Sugar()
	return &Logger{sugar}, nil
}

// Write implements io.Writer interface untuk integration dengan frameworks
func (l *Logger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

// ensureDir membuat directory jika belum ada
func ensureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

// Cleanup method untuk graceful shutdown
func (l *Logger) Sync() error {
	return l.SugaredLogger.Sync()
}
