package main

import (
	routers "web/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// CRUD APIs

	// using data structer we will create Book managment system
	// 1- The there are authors
	// 2- all the authors chould have many books

	// we need to create two struct to have a data between all of them

	// we need to create CRUD APIs for author under api/v1 group
	// Author create author, get author by id, update author, delete author, get all author,
	// Book API, create book, get books by author,  get all books, delete book, update book get all book.
	// We need to create service layer For book and author

	// init gin
	router := gin.Default()
	routers.RegisterTodoRoute(router)
	router.Run()
}
