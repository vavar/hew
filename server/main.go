package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

var DB *Database

var restaurantService *RestaurantService
var userService *UserService

func main() {
	env := os.Getenv("ENVIRONMENT")
	viper.SetConfigName(env)
	viper.AddConfigPath("config")

	Err := viper.ReadInConfig()
	if Err != nil {
		panic(fmt.Errorf("Failed to read the config file: %s", Err))
	}

	DB = InitDBConnection()
	userService := newUserService(DB)
	restaurantService := newRestuarantService(DB)

	router := gin.Default()
	api := router.Group("/api")
	{
	}

	user := api.Group("/user")
	{
		user.GET("/profile", userService.GetProfile)
	}

	restaurant := api.Group("/restuarant")
	{
		restaurant.GET("/info", restaurantService.GetByID)
	}

	admin := api.Group("/admin")
	{
		admin.GET("/users", userService.ListUsers)
		admin.GET("/restuarants", restaurantService.ListRestaurants)
	}

	router.Run() // listen and serve on 0.0.0.0:80
}
