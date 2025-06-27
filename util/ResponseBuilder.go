package responseBuilder

import (
	"fmt"
	response "web/dto/response"
)

func BuildSuccessResponse(obj any, status int, message string) response.GenericResponse {

	response := &response.GenericResponse{}

	fmt.Print("body here")
	fmt.Println(obj)
	response.SetBody(obj)
	response.SetMessage(message)
	response.SetStatus(status)
	response.SetIsSuccess(true)

	fmt.Println(&response)
	fmt.Println(*response)

	return *response
}

func BuildErrorResponse(obj any, status int, message string) response.GenericResponse {

	response := &response.GenericResponse{}

	response.SetBody(obj)
	response.SetMessage(message)
	response.SetStatus(status)

	response.SetIsSuccess(false)

	return *response
}
