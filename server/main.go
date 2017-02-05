package main

import (
	"fmt"
	"os"
	"time"

	"log"

	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	env := os.Getenv("PROFILE")
	viper.SetConfigName(env)
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read the config file: %s", err))
	}

	var db = InitDBConnection()

	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods("DELETE")
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("*")

	router := gin.Default()
	router.Use(cors.New(corsConfig))

	// the jwt middleware
	authMiddleware := &GinJWTMiddleware{
		Realm:      "Hew Platform zone",
		Key:        []byte("Hew leaw tae mai ru ja gin arai"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			var user User
			if err := db.FindUserByEmail(&user, username); err != nil {
				log.Printf("cannot find user by email : %s , %s", username, err)
				return username, false
			}

			if user.Password != password {
				log.Printf("password does not matched")
				return username, false
			}
			return username, true
		},
		Authorizator: func(username string, c *gin.Context) bool {
			if err := db.FindUserByEmail(&User{}, username); err != nil {
				return false
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "invalid.credentials",
				"msg":    "Invalid Credentials",
			})
		},
		Authorized: func(c *gin.Context, username, token string, expire string) {
			var user User
			if db.FindUserByEmail(&user, username) != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "error",
					"error":  "query Failed",
					"msg":    "query Failed",
				})
				return
			}
			c.Writer.Header().Set("token", token)
			c.Writer.Header().Set("Authorization", token)
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data":   user,
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

	var userService = NewUserService(db)

	auth := router.Group("/auth")
	// auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("/refresh", authMiddleware.RefreshHandler)
	auth.GET("/login", authMiddleware.LoginHandler)
	auth.GET("/user", authMiddleware.UserHandler)
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.POST("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
			"msg":    "Logged out Successfully.",
		})
	})
	auth.POST("/register", userService.RegisterUser)
	// auth.GET("/user", userService.GetUserDetails)

	api := router.Group("/api")
	api.GET("/users", userService.ListUsers)
	api.POST("/users", userService.AddUser)
	api.GET("/users/:id", userService.GetByID)
	api.GET("/users/:id/history", userService.ListAllOrders)
	api.PUT("/users", userService.UpdateUser)
	//cross endpoint

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
	api.DELETE("/activities/restaurants", activityService.RemoveRestaurant)
	api.GET("/home", activityService.getOpenOrganizationActivities)

	api.GET("/users/:id/orders", activityService.ListOrderItems)
	api.GET("/orders/info/:id", activityService.GetOrderInfo)
	api.POST("/orders", activityService.CreateOrderItem)
	api.PUT("/orders", activityService.UpdateOrderItem)

	router.Run() // listen and serve on 0.0.0.0:80
}
