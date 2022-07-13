package main

import (
	"gin-rest/src"
	"gin-rest/src/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = ":5000"

func main() {
	db.DbInit()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	src.DefineRoutes(r)
	r.Run(port)
}