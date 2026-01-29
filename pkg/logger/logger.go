package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

// InitLogger initializes the logger
func InitLogger() {
	log = logrus.New()

	// Set log format
	if os.Getenv("LOG_FORMAT") == "json" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
		})
	}

	// Set log level
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warning", "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// Set output
	output := os.Getenv("LOG_OUTPUT")
	logFile := os.Getenv("LOG_FILE")

	if logFile == "" {
		logFile = "logs/app.log"
	}

	// Create logs directory if not exists
	logDir := filepath.Dir(logFile)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Warnf("Failed to create log directory: %v", err)
	}

	// Configure log rotation
	fileWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	switch output {
	case "file":
		log.SetOutput(fileWriter)
	case "both":
		log.SetOutput(os.Stdout)
		log.AddHook(&fileHook{writer: fileWriter})
	default:
		log.SetOutput(os.Stdout)
	}
}

// fileHook writes logs to file
type fileHook struct {
	writer *lumberjack.Logger
}

func (h *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *fileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = h.writer.Write([]byte(line))
	return err
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}

// Helper functions
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return GetLogger().WithFields(fields)
}

// Request logging helper
func LogRequest(method, path, ip string, statusCode int, latency float64) {
	GetLogger().WithFields(logrus.Fields{
		"method":      method,
		"path":        path,
		"ip":          ip,
		"status_code": statusCode,
		"latency_ms":  fmt.Sprintf("%.2f", latency),
	}).Info("HTTP Request")
}
