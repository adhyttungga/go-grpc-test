package config

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DBInit() (*gorm.DB, error) {
	// Connection string for the default database (e.g., "postgres")
	dsnDefault := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", Config.DB.Host, Config.DB.Port, Config.DB.User, Config.DB.Password)

	// Connect to the default database
	dbDefault, err := sql.Open("pgx", dsnDefault) // Use "pgx" driver for PostgreSQL
	if err != nil {
		log.Printf("failed to connect to default database: %v", err)
		return nil, err
	}
	defer dbDefault.Close()

	// Check if the target database exists
	var exists bool
	if err := dbDefault.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", Config.DB.Name).Scan(&exists); err != nil {
		log.Printf("failed to check database existence: %v", err)
		return nil, err
	}

	// Create target database if not exists
	if !exists {
		log.Printf("Database %s does not exist. Creating...\n", Config.DB.Name)
		_, err = dbDefault.Exec(fmt.Sprintf("CREATE DATABASE %q", Config.DB.Name))
		if err != nil {
			log.Printf("failed to create database %s; %v", Config.DB.Name, err)
			return nil, err
		}

		log.Printf("Database '%s' created successfully.\n", Config.DB.Name)
	} else {
		log.Printf("Database '%s' already exists.\n", Config.DB.Name)
	}

	// Connect GORM to the target database
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsnTarget := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Config.DB.Host, Config.DB.Port, Config.DB.User, Config.DB.Password, Config.DB.Name)
	dialect = postgres.Open(dsnTarget)
	dbTarget, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		log.Printf("failed to connect to target database '%s': %v", Config.DB.Name, err)
		return nil, err
	}

	log.Println("Successfully connected to GORM with target database")
	return dbTarget, nil
}
