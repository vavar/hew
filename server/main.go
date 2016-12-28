package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

func errorJSON(c *gin.Context, statusCode int, err error) {
	c.Error(err)
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"error":      err.Error(),
	})
	c.Abort()
}

func main() {
	env := os.Getenv("ENVIRONMENT")
	viper.SetConfigName(env)
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read the config file: %s", err))
	}

	var db = InitDBConnection()

	router := gin.Default()
	api := router.Group("/api")

	var userService = NewUserService(db)
	api.GET("/users", userService.ListUsers)
	api.POST("/users", userService.AddUser)
	api.GET("/users/:id", userService.GetByID)
	api.PUT("/users", userService.UpdateUser)

	var organizationService = NewOrganizationService(db)
	api.GET("/organizations", organizationService.ListOrganizations)
	api.GET("/organizations/:id", organizationService.GetByID)
	api.POST("/organizations", organizationService.CreateOrganization)
	api.PUT("/organizations", organizationService.UpdateOrganization)

	var restaurantService = NewRestaurantService(db)
	api.GET("/restaurants", restaurantService.ListRestaurants)
	api.POST("/restaurants", restaurantService.CreateRestaurant)
	api.PUT("/restaurants", restaurantService.UpdateRestaurant)
	api.GET("/restaurants/:id", restaurantService.GetByID)

	api.POST("/menus", restaurantService.CreateMenu)
	api.PUT("/menus", restaurantService.UpdateMenu)

	var activityService = NewActivityService(db)
	api.GET("/activities/:organizationID", activityService.ListActivities)
	api.POST("/activities", activityService.CreateActivity)
	api.PUT("/activities", activityService.UpdateActivity)

	api.GET("/orders/:userID", activityService.ListOrderItems)
	api.POST("/orders", activityService.CreateOrderItem)
	api.PUT("/orders", activityService.UpdateOrderItem)

	router.Run() // listen and serve on 0.0.0.0:80
}
