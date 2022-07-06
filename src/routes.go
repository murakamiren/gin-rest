package src

import (
	"gin-rest/src/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func testGetV1() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "v1 api test",
		})
	}
}

func DefineRoutes(router gin.IRouter) {
	v1 := router.Group("/v1")
	{
		v1.GET("/test", testGetV1())
		v1.GET("/book", handler.GetAllBook())
	}
}