package handlers

import (
	"classesAndBooking/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookClassInput struct {
	Name string `json:"name" binding:"required"`
	Date string `json:"date" binding:"required"`
}

func BookClass(c *gin.Context) {
	var input BookClassInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, Date : YYYY-MM-DD"})
		return
	}

	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name, Enter a Proper Name with at least 1 characters"})
		return
	}

	booking := models.Booking{
		Name: input.Name,
		Date: date,
	}

	models.Bookings = append(models.Bookings, booking)
	c.JSON(http.StatusOK, gin.H{
		"message": "Bookings created successfully",
		"class":   booking,
	})
}

func GetAllBookings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": models.Bookings})
}
