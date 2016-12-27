package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//UserService - User Service
type UserService struct {
	DB *Database
}

//NewUserService - Instantiate User Service
func NewUserService(db *Database) *UserService {
	return &UserService{db}
}

//GetByID - User Profile by ID
func (service *UserService) GetByID(c *gin.Context) {
	org := &Organization{1, "test"}
	jsonObject := &User{1, "chue", "a@a.com", org}
	c.JSON(http.StatusOK, jsonObject)
}

//AddUser - User Profile by ID
func (service *UserService) AddUser(c *gin.Context) {
	org := &Organization{1, "test"}
	jsonObject := &User{1, "chue", "a@a.com", org}
	c.JSON(http.StatusOK, jsonObject)
}

//ListUsers - list users in System
func (service *UserService) ListUsers(c *gin.Context) {
	org := &Organization{1, "test"}
	jsonObject := struct {
		Users []*User `json:"users"`
	}{[]*User{&User{1, "Chue", "a@a.com", org}, &User{2, "Chue 2", "a@a.com", org}}}
	c.JSON(http.StatusOK, jsonObject)
}
