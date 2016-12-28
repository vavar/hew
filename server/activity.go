package main

import (
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

//ActivityService - activity service
type ActivityService struct {
	DB *Database
}

//NewActivityService - Instantiate Activity Service
func NewActivityService(db *Database) *ActivityService {
	return &ActivityService{db}
}

//ListActivities - list related activity
func (service *ActivityService) ListActivities(c *gin.Context) {
	organizationID, err := strconv.Atoi(c.Param("organizationID"))
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var activities []Activity
	if err := service.DB.ListActivities(&activities, organizationID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, activities)
}

//CreateActivity - create activity
func (service *ActivityService) CreateActivity(c *gin.Context) {
	var json Activity
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Create(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var activity Activity
	if err := service.DB.FindActivityByID(&activity, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, activity)
}

//UpdateActivity - update activity
func (service *ActivityService) UpdateActivity(c *gin.Context) {
	var json Activity
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Update(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var activity Activity
	if err := service.DB.FindActivityByID(&activity, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, activity)
}

//ListOrderItems - list order item
func (service *ActivityService) ListOrderItems(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var orderItems []OrderItem
	if err := service.DB.ListOrderItems(&orderItems, userID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, orderItems)
}

//CreateOrderItem - create order item
func (service *ActivityService) CreateOrderItem(c *gin.Context) {
	var json OrderItem
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Create(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var orderItem OrderItem
	if err := service.DB.FindOrderItemByID(&orderItem, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, orderItem)
}

//UpdateOrderItem - update order item
func (service *ActivityService) UpdateOrderItem(c *gin.Context) {
	var json OrderItem
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Update(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var orderItem OrderItem
	if err := service.DB.FindOrderItemByID(&orderItem, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, orderItem)
}
