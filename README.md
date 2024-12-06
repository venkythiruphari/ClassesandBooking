# Classes and Booking API

This project provides a simple API for creating and booking classes using the Gin framework in Go.

## Prerequisites

1. Install Go

## Setup

1. Clone the repository:

  git clone https://github.com/venkythiruphari/ClassesandBooking.git

2. Navigate to the project directory:

  cd ClassesandBooking

3. Install dependencies:

  go mod tidy

4. Run the application:

  go run main.go

5. Run the simple unit testcases:

  go test ./handlers -v `or` go test ./handlers

6. The API will be available at http://localhost:6000.

### API Endpoints

##### Create Class
URL: `/classes` - `http://localhost:6000/classes`
Method: POST
#### Request Body:

{
  "name": "Pilates",
  "start_date": "2024-12-05",
  "end_date": "2024-12-05",
  "capacity": 10
}

#### Success Response:
Code: 200 OK
Content:

{
  "class": {
    "Name": "Pilates",
    "StartDate": "2024-12-05T00:00:00Z",
    "EndDate": "2024-12-05T00:00:00Z",
    "Capacity": 10
  },
  "message": "Class created successfully"
}

### API Endpoints

##### Book Class
URL: `/bookings` - `http://localhost:6000/bookings`
Method: POST
Request Body:

{
  "name":"Venkanna",
  "date": "2024-12-05"
}

#### Success Response:
Code: 200 OK
Content:

{
  "class": {
    "Name": "Venkanna",
    "Date": "2024-12-05T00:00:00Z"
  },
  "message": "Class created successfully"
}

This project uses an in-memory array to store data. Data will be lost when the server restarts.
Ensure proper date formats (YYYY-MM-DD) for start_date, end_date, and date.
