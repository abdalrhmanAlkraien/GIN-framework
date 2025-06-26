package todoMiddleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TodoMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		fmt.Printf("[%s] [%s] [%s] \n", ctx.RemoteIP(), ctx.Request.RequestURI, ctx.Request.Method)
		ctx.Next()
		fmt.Println("request end")
	}
}
