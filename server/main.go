package main

import (
	"fmt"
	"os"
	"time"

	"./jwt"
	"github.com/spf13/viper"
	cors "gopkg.in/gin-contrib/cors.v1"
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
	router.Use(cors.Default())

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "Hew Platform zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
	}

	router.POST("/login", authMiddleware.LoginHandler)
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"text": "Hello World.",
			})
		})
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

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
	api.GET("/activities", activityService.ListActivities)
	api.POST("/activities", activityService.CreateActivity)
	api.PUT("/activities", activityService.UpdateActivity)
	api.POST("/activities/restaurants", activityService.AddRestaurant)

	api.GET("/orders/:userID", activityService.ListOrderItems)
	api.POST("/orders", activityService.CreateOrderItem)
	api.PUT("/orders", activityService.UpdateOrderItem)

	router.Run() // listen and serve on 0.0.0.0:80
}
