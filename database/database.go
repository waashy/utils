package database

import (
	config "github.com/waashy/utils/database/config"
	dberror "github.com/waashy/utils/database/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Config   config.DBConfig
	Instance *gorm.DB
}

func NewDatabase(config config.DBConfig) (*Database, error) {
	dsn := config.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, dberror.ErrDbConn
	}

	return &Database{
		Config:   config,
		Instance: db,
	}, nil
}
