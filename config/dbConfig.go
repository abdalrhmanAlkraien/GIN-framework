package config

import (
	"web/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func DbConfig() {
	dsn := "host=localhost user=admin password=admin dbname=access-control_service_db port=5434 sslmode=disable TimeZone=Asia/Shanghai"
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
