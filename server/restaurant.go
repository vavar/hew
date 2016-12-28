package main

import (
	"net/http"
	"strconv"

	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

//RestaurantService - Manage Restaurant Data
type RestaurantService struct {
	DB *Database
}

//NewRestaurantService - Instantiate Restaurant Service
func NewRestaurantService(db *Database) *RestaurantService {
	return &RestaurantService{db}
}

//ListRestaurants - list Restaurant in System
func (service *RestaurantService) ListRestaurants(c *gin.Context) {
	var restaurants []Restaurant
	err := service.DB.ListRestaurants(&restaurants)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

//CreateRestaurant - add new Restaurant
func (service *RestaurantService) CreateRestaurant(c *gin.Context) {
	var restJSON Restaurant
	if bindErr := c.BindJSON(&restJSON); bindErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": bindErr})
		return
	}
	if err := service.DB.CreateRestaurant(&restJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, restJSON)
}

//UpdateRestaurant - update exists Restaurant
func (service *RestaurantService) UpdateRestaurant(c *gin.Context) {
	var restJSON Restaurant
	if bindErr := c.BindJSON(&restJSON); bindErr != nil || restJSON.ID <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": bindErr})
		return
	}
	if err := service.DB.UpdateRestaurant(&restJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, restJSON)
}

//CreateMenu - add new Menu
func (service *RestaurantService) CreateMenu(c *gin.Context) {
	var menuJSON Menu
	if bindErr := c.BindJSON(&menuJSON); bindErr != nil {
		log.Printf("bind JSON failed : %s ", bindErr)
		c.JSON(http.StatusInternalServerError, gin.H{"message": bindErr})
		return
	}
	if err := service.DB.Create(&menuJSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, menuJSON)
}

//UpdateMenu - update exists Menu
func (service *RestaurantService) UpdateMenu(c *gin.Context) {
	var json *Menu
	if c.BindJSON(&json) == nil {
		if service.DB.UpdateMenu(json) == nil {
			c.JSON(http.StatusOK, json)
		}
	}

}

//GetByID - get Restaurant by ID
func (service *RestaurantService) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var restaurant Restaurant
	err = service.DB.FindRestaurantByID(&restaurant, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}
