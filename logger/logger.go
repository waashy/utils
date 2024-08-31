package logger

import (
	log "log/slog"
	"os"
)

const DEBUG = "DEBUG"
const ERROR = "ERROR"
const WARN = "WARN"
const INFO = "INFO"

func NewLogConfig(logLevel string) *log.JSONHandler {
	return log.NewJSONHandler(os.Stdout, &log.HandlerOptions{
		Level: getLogLevel(logLevel),
	})
}

func getLogLevel(logLevel string) log.Leveler {
	switch logLevel {
	case DEBUG:
		return log.LevelDebug
	case ERROR:
		return log.LevelError
	case WARN:
		return log.LevelWarn
	case INFO:
		return log.LevelInfo
	default:
		return log.LevelInfo
	}
}

func NewLogger(logLevel string) (*log.Logger, error) {
	logger := log.New(NewLogConfig(logLevel))
	return logger, nil
}
