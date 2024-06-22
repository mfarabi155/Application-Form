package controllers

import (
	"application-form/database"
	"application-form/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreatePersonInput struct {
	Name           string `json:"name" binding:"required"`
	IdentityNumber string `json:"identity_number" binding:"required"`
	Email          string `json:"email" binding:"required"`
	DateOfBirth    string `json:"date_of_birth" binding:"required"`
}

type UpdatePersonInput struct {
	Name           string `json:"name"`
	IdentityNumber string `json:"identity_number"`
	Email          string `json:"email"`
	DateOfBirth    string `json:"date_of_birth"`
}

func CreatePerson(c *gin.Context) {
	var input CreatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateOfBirth, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	person := models.Person{
		Name:           input.Name,
		IdentityNumber: input.IdentityNumber,
		Email:          input.Email,
		DateOfBirth:    dateOfBirth,
	}

	if err := database.DB.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func FindPersons(c *gin.Context) {
	var persons []models.Person
	if err := database.DB.Find(&persons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func FindPersonByIdentityNumber(c *gin.Context) {
	var person models.Person
	identityNumber := c.Param("identity_number")

	if err := database.DB.Where("identity_number = ?", identityNumber).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found!"})
		return
	}

	c.JSON(http.StatusOK, person)
}

func UpdatePersonByIdentityNumber(c *gin.Context) {
	var person models.Person
	identityNumber := c.Param("identity_number")

	if err := database.DB.Where("identity_number = ?", identityNumber).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found!"})
		return
	}

	var input UpdatePersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != "" {
		person.Name = input.Name
	}
	if input.IdentityNumber != "" {
		person.IdentityNumber = input.IdentityNumber
	}
	if input.Email != "" {
		person.Email = input.Email
	}
	if input.DateOfBirth != "" {
		dateOfBirth, err := time.Parse("2006-01-02", input.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
		person.DateOfBirth = dateOfBirth
	}

	if err := database.DB.Save(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func DeletePersonByIdentityNumber(c *gin.Context) {
	var person models.Person
	identityNumber := c.Param("identity_number")

	if err := database.DB.Where("identity_number = ?", identityNumber).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found!"})
		return
	}

	if err := database.DB.Delete(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
