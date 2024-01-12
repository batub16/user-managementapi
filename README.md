# user-managementapi

This project is a simple User Management API implemented in Go using the Gin web framework and SQLite as the database. It provides basic CRUD (Create, Read, Update, Delete) operations for managing user information.

## Project Structure

- **main.go**: The main entry point of the application. It establishes a connection to the SQLite database, sets up the Gin router, and runs the server on port 8080.

- **routers/user_router.go**: Defines the user-related routes and uses the CORS middleware for handling cross-origin resource sharing. The routes include creating, retrieving, updating, and deleting user information.

- **controllers/user_controller.go**: Implements the business logic for handling user operations. It includes functions for getting all users, getting a specific user by ID, creating a new user, updating an existing user, and deleting a user.

## Getting Started
Set up the SQLite database:

Create a SQLite database file (e.g., UserInfo.db) and update the database connection string in main.go.
Run the application:

For Start
go run main.go
The server will be running at http://localhost:8080.

API Endpoints
GET /users: Get a list of all users.
GET /users/:id: Get details of a specific user by ID.
POST /users: Create a new user.
PUT /users/:id: Update an existing user by ID.
DELETE /users/:id: Delete a user by ID.

Dependencies
Gin: Web framework for Go.
Gin CORS Middleware: Middleware for handling Cross-Origin Resource Sharing.
SQLite3 Driver for Go (mattn/go-sqlite3): SQLite3 driver for Go's database/sql package.
