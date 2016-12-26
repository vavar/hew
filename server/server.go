package main

import (
	"fmt"
	"os"
	"github.com/gin-contrib/static"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	viper.SetConfigName(env)
	viper.AddConfigPath("server/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read the config file: %s\n", err))
	}

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

	router.Run(":80") // listen and serve on 0.0.0.0:80
}
