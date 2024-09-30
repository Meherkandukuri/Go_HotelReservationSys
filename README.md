# Go Hotel Reservation System

This project is a Hotel Reservation System built using Go (Golang). It provides functionalities for both guests and administrators to manage room bookings, reservations, and more.

## Features
- Room Booking
- Reservation Management
- User Authentication (Guests and Admins)

## Technologies
- Backend: Go (Golang)
- Database: PostgreSQL
- Frontend: HTML, CSS, JavaScript
- Templates: Go templates

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Meherkandukuri/Go_HotelReservationSys.git
   cd Go_HotelReservationSys


2. Install the required dependencies:
    ```bash
    go mod tidy
  

### Running the Application

To start the application, run the following command:

```bash
go run main.go
```
## Directory Structure

- `cmd/web`: Main entry point of the application.
- `internal`: Application-specific logic (handlers, models).
- `migrations`: Database migration files.
- `templates`: HTML templates for rendering views.
- `static`: CSS, JavaScript files.
