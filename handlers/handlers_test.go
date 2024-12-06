package handlers

import (
	"bytes"
	"classesAndBooking/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateClass(t *testing.T) {
	// Reset the data before each test
	models.Classes = []models.Class{}

	router := SetUpRouter()
	router.POST("/classes", CreateClass)

	// Create a request body
	var jsonStr = []byte(`{"name":"Pilates","start_date":"2024-12-01","end_date":"2024-12-20","capacity":10}`)
	req, _ := http.NewRequest("POST", "/classes", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Pilates")

	// Validate that the class was added
	assert.Equal(t, 1, len(models.Classes))
	assert.Equal(t, "Pilates", models.Classes[0].Name)
	assert.Equal(t, 10, models.Classes[0].Capacity)
}

func TestGetAllClasses(t *testing.T) {
	// Reset the data before each test
	models.Classes = []models.Class{
		{
			Name:      "Pilates",
			StartDate: time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 20, 0, 0, 0, 0, time.UTC),
			Capacity:  10,
		},
	}

	router := SetUpRouter()
	router.GET("/classes", GetAllClasses)

	req, _ := http.NewRequest("GET", "/classes", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Pilates")
}

func TestBookClass(t *testing.T) {
	// Reset the data before each test
	models.Bookings = []models.Booking{}

	router := SetUpRouter()
	router.POST("/bookings", BookClass)

	var jsonStr = []byte(`{"name":"Venkanna","date":"2024-12-01"}`)
	req, _ := http.NewRequest("POST", "/bookings", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Venkanna")

	// Validate that the booking was added
	assert.Equal(t, 1, len(models.Bookings))
	assert.Equal(t, "Venkanna", models.Bookings[0].Name)
}

func TestGetAllBookings(t *testing.T) {
	// Reset the data before each test
	models.Bookings = []models.Booking{
		{
			Name: "Venkanna",
			Date: time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	router := SetUpRouter()
	router.GET("/bookings", GetAllBookings)

	req, _ := http.NewRequest("GET", "/bookings", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Venkanna")
}
