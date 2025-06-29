package todoservice

import (
	"errors"
	todo "web/dto/request"
	repository "web/repository"
)

var todos []todo.Todo

func GetTodos() []todo.Todo {

	return todos
}

func CreateNewTodo(request todo.Todo) {

	if len(todos) == 0 {
		request.Id = 1 // Mock ID generation
	} else {
		request.Id = todos[len(todos)-1].Id + 1 // Mock ID generation
	}

	repository.CreateTodoRepository(request)
	todos = append(todos, request)
}

func GetTodoById(id int) (*todo.Todo, error) {

	for _, t := range todos {
		if t.Id == id {
			return &t, nil
		}
	}

	return nil, errors.New("todo is not there")
}
