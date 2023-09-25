# RESTful CRUD API written in Golang
This is a RESTful API written in Go using the Gin framework. It provides endpoints to fulfill the basic CRUD commands (create, read, update and delete) for working with to-do items.

## Getting Started
### Prerequisites

Before running this project, make sure you have the following installed on your system:

- [Go](https://golang.org/doc/install) 

### Installation
1.  Clone the repository: `git clone https://github.com/kimhwany/RESTfulAPI.git`
2.  Set up dependency tracking:  `go mod download` 
3.  Download gin onto local go application: `go get github.com/gin-gonic/gin`
4.  Run the application: `go run main.go`

## Endpoints

- `GET /todos` - Get a list of all todo tasks.
- `GET /todos/:id` - Get a specific todo task by ID.
- `POST /todos` - Create a new todo task.
- `PATCH /todos/:id` - Toggle the status of a specific todo task.
- `DELETE /todos/:id` - Delete a specific todo task.

## Features

- **Gin Framework**: This project uses the Gin web framework for building the API, known for its performance and simplicity in Go.

- **CRUD Operations**: The API supports basic CRUD (Create, Read, Update, Delete) operations for managing todo items:

  - **Create (C)**: Add a new todo task by sending a POST request with JSON data to `/todos`.

  - **Read (R)**: Retrieve all todos with a GET request to `/todos`, or fetch a specific task by ID with `GET /todos/:id`.

  - **Update (U)**: Toggle the status (completed or not completed) of a task using a PATCH request to `/todos/:id`.

  - **Delete (D)**: Remove a task by sending a DELETE request to `/todos/:id`.

### Testing with Postman

You can use [Postman](https://www.postman.com/) to easily test the API endpoints.






