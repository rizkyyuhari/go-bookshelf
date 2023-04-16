package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkyyuhari/go-bookshelf.git/db"
	"github.com/rizkyyuhari/go-bookshelf.git/repository"
)

func main() {
	g := gin.Default()
	db.OpenDatabase()


	//books root
	g.GET("/books",repository.GetAllBook)
	g.POST("/books",repository.CreateBook)
	g.DELETE("/books/:id",repository.DeleteBook)
	g.PATCH("/books/:id",repository.UpdateBook)
	g.Run()

}

