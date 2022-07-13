package handler

import (
	"gin-rest/src/db"
	"gin-rest/src/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate input
		title := c.PostForm("title")
		author := c.PostForm("author")

		// insert
		book := types.Book{Title: title ,Author: author}
		println(&book)
		result := db.DB.Create(&book)

		if result.Error != nil {
			panic(result.Error)
		}

		c.JSON(http.StatusOK, gin.H{"data": book})
	}
}

func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "nice book",
			"author": "fujie",
		})
	}
}

// func GetAllBooks() gin.HandlerFunc {
// 	return func (c *gin.Context) {
// 		c.JSON(http.StatusOK, books)
// 	}
// }