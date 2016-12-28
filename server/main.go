package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

//DB is an instance of Database type
var database *Database

var activityService *ActivityService
var organizationService *OrganizationService
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

	database = InitDBConnection()
	defer database.DB.Close()
	activityService = NewActivityService(database)
	organizationService = NewOrganizationService(database)
	restaurantService = NewRestaurantService(database)
	userService = NewUserService(database)

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/users", userService.ListUsers)
	api.POST("/users", userService.AddUser)
	api.GET("/users/:id", userService.GetByID)
	api.PUT("/users", userService.UpdateUser)

	api.GET("/organizations", organizationService.ListOrganizations)
	api.GET("/organizations/:id", organizationService.GetByID)
	api.POST("/organizations", organizationService.CreateOrganization)
	api.PUT("/organizations", organizationService.UpdateOrganization)

	api.GET("/restaurants", restaurantService.ListRestaurants)
	api.POST("/restaurants", restaurantService.CreateRestaurant)
	api.PUT("/restaurants", restaurantService.UpdateRestaurant)
	api.GET("/restaurants/:id", restaurantService.GetByID)

	api.POST("/menus", restaurantService.CreateMenu)
	api.PUT("/menus", restaurantService.UpdateMenu)

	api.GET("/activities/:organizationID", activityService.ListActivities)
	api.POST("/activities", activityService.CreateActivity)
	api.PUT("/activities", activityService.UpdateActivity)

	api.GET("/orders/:userID", activityService.ListOrderItems)
	api.POST("/orders", activityService.CreateOrderItem)
	api.PUT("/orders", activityService.UpdateOrderItem)

	router.Run() // listen and serve on 0.0.0.0:80
}
