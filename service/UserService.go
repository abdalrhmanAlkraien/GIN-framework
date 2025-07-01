package service

import (
	request "web/dto/request"
	response "web/dto/response"
	repository "web/repository"
)

func CreateUser(user request.UserRequest) error {

	return repository.CreateUser(user)
}

func GetUsers() ([]*response.UserResponse, error) {

	return repository.GetUsers()
}

func GetUserById(id int) (*response.UserResponse, error) {

	return repository.GetUserById(id)
}
