package route

import (
	controller "web/controller"
	userMiddleware "web/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(r *gin.Engine) {

	userGroup := r.Group("/users")

	userGroup.Use(userMiddleware.UserMiddleware())
	{
		userGroup.POST("/", CreateUser)
		userGroup.GET("/", GetAllUser)
		userGroup.GET("/:id", GetUserById)
	}
}

func CreateUser(c *gin.Context) {

	controller.CreateUser(c)
}

func GetAllUser(c *gin.Context) {

	controller.GetAllUsers(c)
}

func GetUserById(c *gin.Context) {

	controller.GetUserById(c)
}
