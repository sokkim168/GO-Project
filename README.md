# GO-Room-Booking-Project

This is a simple room booking project built with Golang and GORM. The project demonstrates how to manage room and room booking records, including creating, retrieving, updating, and deleting records.

## Getting Started

To get started with this project, follow the instructions below.

### Prerequisites

- Go (version 1.16 or later)
- MySQL database

### Setup

1. **Clone the repository**:

    ```sh
    git clone https://github.com/sokkim168/GO-Project.git
    cd GO-Project
    ```

2. **Install dependencies**:

    ```sh
    go mod tidy
    ```

3. **Configure the database**:

    Update the database connection details in your `.env` on the root directory.

4. **Run migrations**:

    This is simple, you just routing localhost:8080/api/rooms or localhost:8080/api/room-booking with get method then it will auto create a database name room_booking if doesn't exsit and then will migrations all those table 

### Running the Project

To build and run the project, use the following commands:

1. **Build the project**:

    ```sh
    go build
    ```

    This command will build the project and resolve dependencies.

2. **Run the project**:

    ```sh
    go run .
    ```

    This command will compile and run the Go code.

### API Endpoints

Here are the available API endpoints:

- **Rooms**:
  - `GET /api/rooms` - Get all rooms
  - `GET /api/rooms/{id:[0-9]+}` - Get a room by ID
  - `POST /api/rooms` - Create a new room
  - `PUT /api/rooms/{id:[0-9]+}` - Update a room by ID
  - `DELETE /api/rooms/{id:[0-9]+}` - Delete a room by ID

- **Room Bookings**:
  - `GET /api/room_bookings` - Get all room bookings
  - `GET /api/room_bookings/{id:[0-9]+}` - Get a room booking by ID
  - `POST /api/room_bookings` - Create a new room booking
  - `PUT /api/room_bookings/{id:[0-9]+}` - Update a room booking by ID
  - `DELETE /api/room_bookings/{id:[0-9]+}` - Delete a room booking by ID

### Example Request

To create a new room, send a POST request to `/api/rooms` with the following JSON body:
```json
{
  "RoomNumber": "101",
  "RoomSize": "Large",
  "Floor": 1,
  "Status": "Available"
}
```
To create a new room booking, send a POST request to `/api/room-bookings` with the following JSON body:
```json
{
    "RoomID": 1,
    "BookingDate": "2024-06-17 12:10:50",
    "StartDate": "2024-06-18 20:10:51",
    "EndDate": "2024-06-20 05:10:52",
    "Status": "VIP"
}
