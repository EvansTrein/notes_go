package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Привет я GIN",
		})
	})

	router.Run(":9100")
}