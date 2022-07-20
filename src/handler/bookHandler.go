package handler

import (
	"gin-rest/src/db"
	"gin-rest/src/types"
	"net/http"
	"strconv"

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

func GetBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book types.Book
		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		db.DB.Where("id = ?", id).Find(&book)
		c.JSON(http.StatusOK, gin.H{"data": book})
	}
}

func GetAllBooks() gin.HandlerFunc {
	return func (c *gin.Context) {
		var books []types.Book
		db.DB.Find(&books)
		c.JSON(http.StatusOK, gin.H{"data": books})
	}
}

func UpdateBook() gin.HandlerFunc {
	return func (c *gin.Context) {
		title := c.PostForm("title")
		author := c.PostForm("author")
		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		db.DB.Model(&types.Book{}).Where("id = ?" ,id).Updates(types.Book{Title: title, Author: author})
	}
}

func DeleteBook() gin.HandlerFunc {
	return func (c *gin.Context) {
		idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)

		db.DB.Delete(&types.Book{}, id)
	}
}