package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var bookData []types.BookStruct = types.BookStruct[
// 	{
// 		title: "nice book",
// 		disc: "this is a nice book",
// 		value: 300,
// 		author: "fujie",
// 		isPublished: false,
// 	}
// ]

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