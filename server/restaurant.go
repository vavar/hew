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
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

//CreateRestaurant - add new Restaurant
func (service *RestaurantService) CreateRestaurant(c *gin.Context) {
	var restJSON Restaurant
	if err := c.BindJSON(&restJSON); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.CreateRestaurant(&restJSON); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, restJSON)
}

//UpdateRestaurant - update exists Restaurant
func (service *RestaurantService) UpdateRestaurant(c *gin.Context) {
	var restJSON Restaurant
	if err := c.BindJSON(&restJSON); err != nil || restJSON.ID <= 0 {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Update(&restJSON); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, restJSON)
}

//CreateMenu - add new Menu
func (service *RestaurantService) CreateMenu(c *gin.Context) {
	var menuJSON Menu
	if err := c.BindJSON(&menuJSON); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Create(&menuJSON); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, menuJSON)
}

//UpdateMenu - update exists Menu
func (service *RestaurantService) UpdateMenu(c *gin.Context) {
	var menuJSON Menu
	if c.BindJSON(&menuJSON) == nil {
		if service.DB.Update(&menuJSON) == nil {
			c.JSON(http.StatusOK, menuJSON)
		}
	}
}

//GetByID - get Restaurant by ID
func (service *RestaurantService) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var restaurant Restaurant
	err = service.DB.FindRestaurantByID(&restaurant, id)
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, restaurant)
}
