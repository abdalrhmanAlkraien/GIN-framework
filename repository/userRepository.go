package repository

import (
	config "web/config"
	model "web/model"
)

func CreateUserRepository(user *model.User) {

	db := config.GetDatabaseConnection()

	db.Create(&user)
}
