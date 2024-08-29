package config

import "fmt"

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	SSLMode  string `json:"ssl_mode"`
}

func (cfg *DBConfig) DSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
}
