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

//ListOrganizations - List all organizations
func (service *OrganizationService) ListOrganizations(c *gin.Context) {
	var organizations []Organization
	if err := service.DB.ListOrganizations(&organizations); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, organizations)
}

//GetByID - Get an organization with the given ID
func (service *OrganizationService) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var organization Organization
	if err = service.DB.FindOrganizationByID(&organization, id); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, organization)
}

//CreateOrganization - Create an organization
func (service *OrganizationService) CreateOrganization(c *gin.Context) {
	var json Organization
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Create(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var organization Organization
	if err := service.DB.FindOrganizationByID(&organization, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, organization)
}

//UpdateOrganization - Update an organization
func (service *OrganizationService) UpdateOrganization(c *gin.Context) {
	var json Organization
	if err := c.BindJSON(&json); err != nil {
		errorJSON(c, http.StatusBadRequest, err)
		return
	}

	if err := service.DB.Update(&json); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	var organization Organization
	if err := service.DB.FindOrganizationByID(&organization, json.ID); err != nil {
		errorJSON(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, organization)
}
