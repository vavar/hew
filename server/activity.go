package main

import (
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

}

//CreateActivity - create activity
func (service *ActivityService) CreateActivity(c *gin.Context) {

}

//UpdateActivity - update activity
func (service *ActivityService) UpdateActivity(c *gin.Context) {

}

//ListOrderItem - list order item
func (service *ActivityService) ListOrderItem(c *gin.Context) {

}

//CreateOrderItem - create order item
func (service *ActivityService) CreateOrderItem(c *gin.Context) {

}

//UpdateOrderItem - update order item
func (service *ActivityService) UpdateOrderItem(c *gin.Context) {

}
