package main

import (
	log "log/slog"
	"os"

	config "github.com/waashy/see-user/app/model/config"
	"github.com/waashy/utils/config/parser"
)

const CONFIG_FILE_PATH = "../config/runConfig.json"

func init() {}

func main() {
	logger := log.New(log.NewJSONHandler(os.Stdout, &log.HandlerOptions{Level: log.LevelDebug}))

	logger.Info("Server Starting")

	var appConfig config.AppConfig

	err := parser.ConfigParser(CONFIG_FILE_PATH, &appConfig)
	if err != nil {
		logger.Error("Failed to parse the configuration", "Err", err)
	}

}
