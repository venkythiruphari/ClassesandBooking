package main

import (
	"classesAndBooking/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	const Port = ":6000"
	router := gin.Default()

	// Class Routes
	router.POST("/classes", handlers.CreateClass)
	router.GET("/classes", handlers.GetAllClasses)

	// Booking Routes
	router.POST("/bookings", handlers.BookClass)
	router.GET("/bookings", handlers.GetAllBookings)

	log.Println("Server is running on port", Port)
	if err := router.Run(Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
