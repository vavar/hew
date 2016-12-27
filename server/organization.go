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
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var organization Organization
	err = service.DB.FindOrganizationByID(&organization, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, organization)
}

//CreateOrganization - Create an organization
func (service *OrganizationService) CreateOrganization(c *gin.Context) {
	var json Organization
	if c.BindJSON(&json) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if service.DB.CreateOrganization(&json) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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

	if err := service.DB.UpdateOrganization(&json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated the organization successfully",
	})
}
