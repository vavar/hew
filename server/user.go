package main

import (
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var user User
	if err = service.DB.FindUserByID(&user, id); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

//AddUser - User Profile by ID
func (service *UserService) AddUser(c *gin.Context) {
	var json User
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Create(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var user User
	if err := service.DB.FindUserByID(&user, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

//ListUsers - list users in System
func (service *UserService) ListUsers(c *gin.Context) {
	var users []User
	if err := service.DB.ListUsers(&users); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

//UpdateUser - Update a user
func (service *UserService) UpdateUser(c *gin.Context) {
	var json User
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Update(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var user User
	if err := service.DB.FindUserByID(&user, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
