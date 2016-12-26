package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong pong pong asdasda ssss ",
			})
		})

		api.GET("/pong", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong pong pong asdasda ssss ",
			})
		})
	}

	router.Run(":80") // listen and serve on 0.0.0.0:80
}
