package config

import "fmt"

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	SSLMode  string
}

func (cfg *DBConfig) DSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode)
}
