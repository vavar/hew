package main

import (
	"net/http"
	"strconv"

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

}

//CreateMenu - add new Menu
func (service *RestaurantService) CreateMenu(c *gin.Context) {

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
