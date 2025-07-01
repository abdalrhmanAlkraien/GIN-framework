package controller

import (
	"fmt"
	"net/http"
	"strconv"
	todo "web/dto/request"
	todoService "web/service"
	response "web/util"

	"github.com/gin-gonic/gin"
)

func CreateTodoController(ctx *gin.Context) {

	var todo todo.Todo // I will fill the data via pointer

	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse(nil, http.StatusBadRequest, err.Error()))
		return // just for exist from function no need the gin here no expected return value.
	}

	err := todoService.CreateNewTodo(todo)

	if err != nil {

		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, response.BuildSuccessResponse(nil, http.StatusCreated, "Todo has been created"))
}

func GetAllTodos(ctx *gin.Context) {

	// var todos []todo.Todo

	todos, error := todoService.GetTodos()

	if error != nil {
		ctx.JSON(http.StatusBadRequest, response.BuildErrorResponse(nil, http.StatusBadRequest, error.Error()))
	}
	ctx.JSON(http.StatusOK, response.BuildSuccessResponse(
		todos, http.StatusOK, "success request"))
}

func GetTodoById(ctx *gin.Context) {

	idParam, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {

		println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))

	}

	var todo *todo.Todo

	todo, err = todoService.GetTodoById(idParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))
	} else {
		ctx.JSON(http.StatusOK, response.BuildSuccessResponse(todo, http.StatusOK, "Success request"))
	}

}
