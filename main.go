package main

import (
	"github.com/gin-gonic/gin"
)

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	r.POST("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, helloResponse{
			Message: "Hello, " + name + "!",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, helloResponse{
			Message: "Hello, World!",
		})
	})

	r.POST("/hi/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, helloResponse{
			Message: "Hi, " + name + "!",
		})
	})

	r.Run("[::]:8080")
}
