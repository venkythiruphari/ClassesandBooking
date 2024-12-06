package handlers

import (
	"classesAndBooking/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateClassInput struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required"`
}

func CreateClass(c *gin.Context) {
	var input CreateClassInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format, Date : YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format, Date : YYYY-MM-DD"})
		return
	}

	if endDate.Before(startDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End date must be after the start date"})
		return
	}

	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name, enter a proper name with at least 1 characters"})
		return
	}

	if input.Capacity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Capacity must be a positive number"})
		return
	}

	class := models.Class{
		Name:      input.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  input.Capacity,
	}

	models.Classes = append(models.Classes, class)
	c.JSON(http.StatusOK, gin.H{
		"message": "Class created successfully",
		"class":   class,
	})
}

func GetAllClasses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": models.Classes})
}
