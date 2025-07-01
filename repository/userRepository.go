package repository

import (
	config "web/config"
	request "web/dto/request"
	response "web/dto/response"
	"web/model"

	"github.com/jinzhu/copier"
)

func CreateUser(user request.UserRequest) error {

	db := config.GetDatabaseConnection()

	var userModel model.User

	copier.Copy(&userModel, user)

	result := db.Create(&userModel)

	return result.Error
}

func GetUsers() ([]*response.UserResponse, error) {

	db := config.GetDatabaseConnection()

	var users []model.User

	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	var usersResponse []*response.UserResponse

	copier.Copy(&usersResponse, users)

	return usersResponse, nil
}

func GetUserById(id int) (*response.UserResponse, error) {

	db := config.GetDatabaseConnection()

	var user model.User

	result := db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	var response response.UserResponse

	copier.Copy(&response, user)

	return &response, nil
}
