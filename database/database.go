package database

import (
	config "github.com/waashy/see-user/database/config"
	dberror "github.com/waashy/see-user/database/errors"
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

func (db *Database) AutoMigrate(structs []interface{}) error {
	for _, eachStruct := range structs {
		if err := db.Instance.AutoMigrate(eachStruct); err != nil {
			return err
		}
	}
	return nil
}
