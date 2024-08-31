package logger

import (
	log "log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
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

func RequestLogger(logger *log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Log the incoming request
		logger.Info("Incoming request",
			log.String("method", c.Method()),
			log.String("path", c.Path()),
			"body", c.Request(),
		)

		// Continue to the next handler
		err := c.Next()

		// Log the response status
		logger.Info("Response",
			log.String("method", c.Method()),
			log.String("path", c.Path()),
			log.Int("status", c.Response().StatusCode()),
			"response", c.Response(),
		)

		return err
	}
}
