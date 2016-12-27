package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

//DB is an instance of Database type
var DB *Database

var restaurantService *RestaurantService
var userService *UserService
var activityService *ActivityService

func main() {
	env := os.Getenv("ENVIRONMENT")
	viper.SetConfigName(env)
	viper.AddConfigPath("config")

	Err := viper.ReadInConfig()
	if Err != nil {
		panic(fmt.Errorf("Failed to read the config file: %s", Err))
	}

	DB = InitDBConnection()
	defer DB.DB.Close()
	userService = NewUserService(DB)
	restaurantService = NewRestaurantService(DB)

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/users", userService.ListUsers)
	api.POST("/users", userService.AddUser)
	api.GET("/users/:id", userService.GetByID)

	api.GET("/restaurants", restaurantService.ListRestaurants)
	api.POST("/restaurants", restaurantService.CreateRestaurant)
	api.PUT("/restaurants", restaurantService.UpdateRestaurant)
	api.GET("/restaurants/:id", restaurantService.GetByID)

	api.POST("/menus", restaurantService.CreateMenu)
	api.PUT("/menus", restaurantService.UpdateMenu)

	api.GET("/activities", activityService.ListActivities)
	api.POST("/activities", activityService.CreateActivity)
	api.PUT("/activities", activityService.UpdateActivity)

	api.GET("/orders", activityService.ListOrderItem)
	api.POST("/orders", activityService.CreateOrderItem)
	api.PUT("/orders", activityService.UpdateOrderItem)

	router.Run() // listen and serve on 0.0.0.0:80
}
