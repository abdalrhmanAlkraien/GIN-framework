package todoroute

import (
	"net/http"
	"strconv"
	todoMiddleware "web/middleware"
	todo "web/model"
	service "web/service"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoute(r *gin.Engine) { //Engine use for register routes, but context use per request

	todoGroup := r.Group("/todo")
	todoGroup.Use(todoMiddleware.TodoMiddleware())
	{
		todoGroup.GET("/", GetAllTodo)
		todoGroup.POST("/", CreateNewTodo)
		todoGroup.GET("/:id", GetTodoById)
	}

}

func GetAllTodo(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, service.GetTodos())
}

func CreateNewTodo(c *gin.Context) {

	var todo todo.Todo // I will fill the data via pointer

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // just for exist from function no need the gin here no expected return value.
	}

	service.CreateNewTodo(todo)
}

func GetTodoById(ctx *gin.Context) {

	idParam, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {

		println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	var todo *todo.Todo

	todo, err = service.GetTodoById(idParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find the Todo"})
	} else {
		ctx.JSON(http.StatusOK, &todo)
	}
}
