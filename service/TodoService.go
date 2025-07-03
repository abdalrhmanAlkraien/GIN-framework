package service

import (
	todo "web/dto/request"
	repository "web/repository"
)

func GetTodos() ([]*todo.Todo, error) {

	return repository.GetAllTodos()
}

func CreateNewTodo(request todo.Todo) error {

	return repository.CreateTodoRepository(request)
}

func GetTodoById(id int) (*todo.Todo, error) {

	return repository.GetTodoById(id)
}

func getTodoByUserId(id int) ([]*todo.Todo, error) {

	return repository.GetTodoByUserId(id)
}
