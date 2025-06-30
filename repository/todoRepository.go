package repository

import (
	"fmt"
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

func GetAllTodos() ([]*request.Todo, error) {

	db := config.GetDatabaseConnection()

	var todos []model.Todo

	result := db.Find(&todos)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, result.Error
	} else {

		var todosResponse []*request.Todo
		copier.Copy(&todosResponse, todos)
		return todosResponse, nil
	}
}

func GetTodoById(id int) (*request.Todo, error) {

	var todo *model.Todo
	db := config.GetDatabaseConnection()
	result := db.First(&todo, id)
	fmt.Println("log result")
	fmt.Println(todo)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, result.Error
	} else {

		var todoResponse request.Todo
		copier.Copy(&todoResponse, todo)
		fmt.Println("log response")
		fmt.Println(todoResponse)
		// fmt.Println(&todoResponse)
		return &todoResponse, nil
	}

}
