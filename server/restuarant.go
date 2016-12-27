package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//RestaurantService - Manage Restaurant Data
type RestaurantService struct {
	DB *Database
}

func newRestuarantService(db *Database) *RestaurantService {
	return &RestaurantService{db}
}

//ListRestaurants - list restuarant in System
func (service *RestaurantService) ListRestaurants(c *gin.Context) {
	jsonObject := struct {
		Restuarants []*Restaurant `json:"restaurants"`
	}{[]*Restaurant{&Restaurant{1, "ครัวคุณวี", []*Menu{{1, "ข้าวไข่เจียว", 100}}}}}
	c.JSON(http.StatusOK, jsonObject)
}

//GetByID - get Restaurant by ID
func (service *RestaurantService) GetByID(c *gin.Context) {
	jsonObject := &Restaurant{1, "ครัวคุณวี", []*Menu{{1, "ข้าวไข่เจียว", 100}}}
	c.JSON(http.StatusOK, jsonObject)
}
