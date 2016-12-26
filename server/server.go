package main

import "gopkg.in/gin-gonic/gin.v1"
import "github.com/gin-contrib/static"

func main() {

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../web/dist/", true)))

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

	router.Run() // listen and serve on 0.0.0.0:8080
}
