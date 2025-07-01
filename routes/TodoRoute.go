package route

import (
	todoController "web/controller"
	todoMiddleware "web/middleware"

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

	todoController.GetAllTodos(ctx)
}

func CreateNewTodo(c *gin.Context) {

	todoController.CreateTodoController(c)
}

func GetTodoById(ctx *gin.Context) {

	todoController.GetTodoById(ctx)
}
