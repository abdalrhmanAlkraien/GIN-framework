package repository

import (
	config "web/config"
	request "web/dto/request"
	model "web/model"

	"github.com/jinzhu/copier"
)

func CreateTodoRepository(todo request.Todo) {

	var todoModel model.Todo

	copier.Copy(&todoModel, &todo) // map the data from request to model
	db := config.GetDatabaseConnection()

	db.Create(&todoModel)
}
