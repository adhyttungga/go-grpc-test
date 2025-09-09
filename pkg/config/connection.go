package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DBInit() (*gorm.DB, error) {
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := ""
	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
