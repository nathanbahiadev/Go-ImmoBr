package database

import (
	"github.com/nathanbahiadev/go-immobr/infrastructure/models"
	"github.com/nathanbahiadev/go-immobr/infrastructure/utils"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() (*gorm.DB, error) {
	env := utils.GetEnvVariable("ENVIRONMENT")

	if env == "PRODUCTION" {
		dsn := "host=postgres user=immobr password=immobr dbname=immobr port=5432"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		dsn := "./data.db"
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(&models.RealState{})
	DB.AutoMigrate(&models.User{})

	return DB, nil
}
