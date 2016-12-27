package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//RestaurantService - Manage Restaurant Data
type RestaurantService struct {
	DB *Database
}

//NewRestuarantService - Instantiate Restaurant Service
func NewRestuarantService(db *Database) *RestaurantService {
	return &RestaurantService{db}
}

//ListRestaurants - list restuarant in System
func (service *RestaurantService) ListRestaurants(c *gin.Context) {

	restaurants, err := service.DB.ListRestaurants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// jsonObject := struct {
	// 	Restuarants []*Restaurant `json:"restaurants"`
	// }{[]*Restaurant{rests}}
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

}

//GetByID - get Restaurant by ID
func (service *RestaurantService) GetByID(c *gin.Context) {
	restaurant, err := service.DB.FindRestaurantByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}
