package model

type ServerConfig struct {
	Port int `json:"port"`
}

type AppConfig struct {
	Server   ServerConfig `json:"server_config"`
	LogLevel string       `json:"log_level"`
}
