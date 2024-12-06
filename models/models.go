package models

import (
	"time"
)

type Class struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

type Booking struct {
	Name string
	Date time.Time
}

var Classes []Class
var Bookings []Booking
