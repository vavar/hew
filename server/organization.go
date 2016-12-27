package main

import (
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

//OrganizationService - Organization Service
type OrganizationService struct {
	DB *Database
}

//NewOrganizationService - Instantiate Organization Service
func NewOrganizationService(db *Database) *OrganizationService {
	return &OrganizationService{db}
}

//GetByID - Get an organization with the given ID
func (service *OrganizationService) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	var organization Organization
	if err = service.DB.FindOrganizationByID(&organization, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, organization)
}

//CreateOrganization - Create an organization
func (service *OrganizationService) CreateOrganization(c *gin.Context) {
	var json Organization
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if err := service.DB.Create(&json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Created the organization successfully",
	})
}

//UpdateOrganization - Update an organization
func (service *OrganizationService) UpdateOrganization(c *gin.Context) {
	var json Organization
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if err := service.DB.Update(&json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated the organization successfully",
	})
}
