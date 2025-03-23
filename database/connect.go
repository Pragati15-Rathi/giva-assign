package database

import (
	"fmt"
	"giva-url-shortner/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		config.Config("DB_HOST"),
		config.Config("DB_USERNAME"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_DATABASE"),
		config.Config("DB_PORT"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&GoShort{})
	if err != nil {
		panic(err)
	}
}
