package service

import (
	todo "web/dto/request"
	repository "web/repository"
)

var todos []todo.Todo

func GetTodos() ([]*todo.Todo, error) {

	return repository.GetAllTodos()
}

func CreateNewTodo(request todo.Todo) {

	repository.CreateTodoRepository(request)
	todos = append(todos, request)
}

func GetTodoById(id int) (*todo.Todo, error) {

	return repository.GetTodoById(id)
}
