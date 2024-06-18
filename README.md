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

    This is simple, you just routing localhost:8080/rooms or localhost:8080/room-booking with get method then it will auto create a database name room_booking if doesn't exsit and then will migrations all those table 

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
  - `GET /rooms` - Get all rooms
  - `GET /rooms/{id:[0-9]+}` - Get a room by ID
  - `POST /rooms` - Create a new room
  - `PUT /rooms/{id:[0-9]+}` - Update a room by ID
  - `DELETE /rooms/{id:[0-9]+}` - Delete a room by ID

- **Room Bookings**:
  - `GET /room_bookings` - Get all room bookings
  - `GET /room_bookings/{id:[0-9]+}` - Get a room booking by ID
  - `POST /room_bookings` - Create a new room booking
  - `PUT /room_bookings/{id:[0-9]+}` - Update a room booking by ID
  - `DELETE /room_bookings/{id:[0-9]+}` - Delete a room booking by ID

### Example Request

To create a new room, send a POST request to `/rooms` with the following JSON body:

```json
{
  "RoomNumber": "101",
  "RoomSize": "Large",
  "Floor": 1,
  "Status": "Available"
}
