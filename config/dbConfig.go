package config

import (
	"fmt"
	"os"
	"web/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func DbConfig() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSLMODE")
	dbTZ := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSL, dbTZ,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Todo{}, &model.User{})
	db.AllowGlobalUpdate = true

	database = db
}

func GetDatabaseConnection() *gorm.DB {

	return database
}
