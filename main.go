package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo represents data about the todo task.
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// todos slice to seed record todo task data.
var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

// getTodos responds with the list of all todos as JSON
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) //convert data structure to JSON
}

// addTodo adds an todo from JSON received in the request body
func addTodo(context *gin.Context) {
	var newTodo todo

	// Call BindJSON to bind the received JSON to newTodo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new todo task to the slice.
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// getTodo responds with the list of all todos as JSON
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

// deleteTodo deletes a todo task by ID
func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index, err := findTodoIndexById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Remove the todo task from the slice. 
	todos = append(todos[:index], todos[index+1:]...)
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// findTodoIndexById finds the index of a todo task by ID
func findTodoIndexById(id string) (int, error) {
	for index, task := range todos {
		if task.ID == id {
			return index, nil
		}
	}
	return -1, errors.New("todo not found")
}


// getTodoById locates the todo task whose ID value matches the id parameter
// sent by the cient, then returns that todo task as a response.
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default() // creating server
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.DELETE("todos/:id", deleteTodo)
	router.Run("localhost:9090") //specifying the server path
}
