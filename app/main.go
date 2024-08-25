package main

import (
	log "log/slog"
	"os"

	"github.com/waashy/utils/config/parser"
)

const CONFIG_FILE_PATH = "../config/runConfig.json"

type ServerConfig struct {
	Port int `json:"port"`
}

type AppConfig struct {
	Server ServerConfig `json:"server_config"`
}

func init() {}

func main() {
	logger := log.New(log.NewJSONHandler(os.Stdout, &log.HandlerOptions{Level: log.LevelDebug}))

	logger.Info("Server Starting")

	var appConfig AppConfig

	err := parser.ConfigParser(CONFIG_FILE_PATH, &appConfig)
	if err != nil {
		logger.Error("Failed to parse the configuration", "Err", err)
	}

}
