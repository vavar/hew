package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//AdminService data type
type AdminService struct {
	DB *Database
}

func newAdminService(db *Database) *AdminService {
	return &AdminService{db}
}

//Users in System
func (admin *AdminService) Users(c *gin.Context) {
	org := &Organization{1, "test"}
	jsonObject := struct {
		Users []*User `json:"users"`
	}{[]*User{&User{1, "Chue", "a@a.com", org}, &User{2, "Chue 2", "a@a.com", org}}}
	c.JSON(http.StatusOK, jsonObject)
}

//Restaurants in System
func (admin *AdminService) Restaurants(c *gin.Context) {
	jsonObject := struct {
		Restuarants []*Restaurant `json:"restaurants"`
	}{[]*Restaurant{&Restaurant{1, "ครัวคุณวี", []*Menu{{1, "ข้าวไข่เจียว", 100}}}}}
	c.JSON(http.StatusOK, jsonObject)
}
