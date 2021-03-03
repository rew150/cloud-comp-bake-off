package main

import (
	"lilandfriends/bakeoff/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.POST("/account", handler.CreateAccount)
	r.POST("/login", handler.Login)
	r.POST("/password", handler.UpdatePassword)

	r.Run(":8080")
}
