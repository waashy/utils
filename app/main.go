package main

import (
	log "log/slog"

	config "github.com/waashy/see-user/app/model/config"
	parser "github.com/waashy/utils/config/parser"
	utilLogger "github.com/waashy/utils/logger"
)

const CONFIG_FILE_PATH = "../config/runConfig.json"

var (
	logger *log.Logger
)

func init() {}

func main() {

	log.Info("See-User service starting")

	var appConfig config.AppConfig

	err := parser.ConfigParser(CONFIG_FILE_PATH, &appConfig)
	if err != nil {
		logger.Error("Failed to parse the configuration", "Err", err)
	}

	logger, err := utilLogger.NewLogger("DEBUG")
	if err != nil {
		log.Error("failed to start server", "ERR", err)
		return
	}

	logger.Info("logger configured")
}
