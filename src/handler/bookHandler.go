package handler

import (
	"gin-rest/src/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "nice book",
			"disc": "this is a nice book",
			"value": 300,
			"author": "fujie",
			"isPublished": false,
		})
	}
}

var books = []types.BookStruct{
	{Title: "nice book",Disc: "this is a nice book", Value: 300, Author: "fujie", IsPublished: false},
	{Title: "bad book",Disc: "this is a bad book", Value: 500, Author: "fujie", IsPublished: true},
	{Title: "golang book",Disc: "this is a golang book for beginners", Value: 1000, Author: "go", IsPublished: true},
}

func GetAllBooks() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, books)
	}
}