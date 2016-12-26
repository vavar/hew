package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

var DB *gorm.DB

func main() {
	env := os.Getenv("ENVIRONMENT")
	viper.SetConfigName(env)
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read the config file: %s\n", err))
	}

	// DB = InitDBConnection()

	router := gin.Default()
	api := router.Group("/api")
	{
	}

	user := api.Group("/user")
	{
		userService := newUserService(DB)
		user.GET("/profile", userService.Profile)
	}

	admin := api.Group("/admin")
	{
		adminService := newAdminService(DB)
		admin.GET("/users", adminService.Users)
		admin.GET("/restuarants", adminService.Restaurants)
	}

	router.Run() // listen and serve on 0.0.0.0:80
}
