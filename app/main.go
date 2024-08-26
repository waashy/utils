package main

import (
	log "log/slog"

	config "github.com/waashy/see-user/app/model/config"
	"github.com/waashy/utils/config/parser"
)

const CONFIG_FILE_PATH = "../config/runConfig.json"

var (
	logger = logger.NewLogger("DEBUG")
)

func init() {}

func main() {

	log.Info("See-User service starting")

	var appConfig config.AppConfig

	err := parser.ConfigParser(CONFIG_FILE_PATH, &appConfig)
	if err != nil {
		logger.Error("Failed to parse the configuration", "Err", err)
	}

	logger := logger.NewLogger("DEBUG")

}
