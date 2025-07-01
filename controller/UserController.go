package controller

import (
	"fmt"
	"net/http"
	"strconv"
	request "web/dto/request"
	service "web/service"
	response "web/util"
	responseBuilder "web/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.BindJSON(&userRequest); err != nil {

		c.JSON(http.StatusBadRequest, responseBuilder.BuildErrorResponse(nil, http.StatusBadRequest, err.Error()))

		return
	}

	err := service.CreateUser(userRequest)
	if err != nil {

		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, responseBuilder.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))
	}
	c.JSON(http.StatusCreated, responseBuilder.BuildSuccessResponse(nil, http.StatusAccepted, "User has been created"))
}

func GetAllUsers(c *gin.Context) {

	users, err := service.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, responseBuilder.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))
	}

	c.JSON(http.StatusOK, responseBuilder.BuildSuccessResponse(
		users, http.StatusOK, "success request"))
}

func GetUserById(c *gin.Context) {

	idParam, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		println(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))

	}

	user, err := service.GetUserById(idParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responseBuilder.BuildErrorResponse(nil, http.StatusInternalServerError, err.Error()))
	}

	c.JSON(http.StatusOK, responseBuilder.BuildSuccessResponse(
		user, http.StatusOK, "success request"))
}
