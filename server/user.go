package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

//AdminService data type
type UserService struct {
	DB *Database
}

func newUserService(db *Database) *UserService {
	return &UserService{db}
}

//Users in System
func (service *UserService) Profile(c *gin.Context) {
	org := &Organization{1, "test"}
	jsonObject := &User{1, "chue", "a@a.com", org}
	c.JSON(http.StatusOK, jsonObject)
}
