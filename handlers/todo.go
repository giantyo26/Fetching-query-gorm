package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Defining a struct type "Todos" to represent a Todo item
type Todos struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `isDone:"isDone"`
}

// Declaring a global variable "todos" as a slice of Todos struct and initializing it with some data
var todos = []Todos{
	{
		Id:     "1",
		Title:  "Cuci Tangan",
		IsDone: true,
	},
	{
		Id:     "2",
		Title:  "Jaga Jarak",
		IsDone: false,
	},
	{
		Id:     "3",
		Title:  "Pakai Masker",
		IsDone: false,
	},
}

// Handler function for finding all todos
func FindTodos(c echo.Context) error {
	c.Response().Header().Set("Content/Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(todos) // Encoding the "todos" slice as JSON and writing it to the response body
}

// Handler function for getting a specific todo by ID
func GetTodo(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	id := c.Param("id") // Getting the value of the "id" parameter from the URL

	var todoData Todos // Declaring a variable to hold the retrieved todo
	var isGetTodo = false

	// Looping through the "todos" slice to find the todo with the given ID
	for _, todo := range todos {
		if id == todo.Id {

			todoData = todo
			isGetTodo = true
		}
	}

	// If the todo with the given ID is not found
	if !isGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found") // Writing an error message to the response body
	}

	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(todoData) // Encoding the retrieved todo as JSON and writing it to the response body

}

// Handler function for creating a new todo
func CreateTodo(c echo.Context) error {
	var data Todos

	// Decoding the request body JSON into the "data" variable
	json.NewDecoder(c.Request().Body).Decode(&data)

	// Appending the new todo data to the "todos" slice
	todos = append(todos, data)

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(todos) //
}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")

	var data Todos
	var IsGetTodo = false

	// Decoding the request body JSON into the "data" variable
	json.NewDecoder(c.Request().Body).Decode(&data)

	for idx, todo := range todos {
		if id == todo.Id {
			todos[idx] = data
			IsGetTodo = true
		}
	}

	if !IsGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(todos)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	var isGetTodo = false
	var index int

	for idx, todo := range todos {
		if id == todo.Id {
			index = idx
			isGetTodo = true
		}
	}

	if !isGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	todos = append(todos[:index], todos[index+1:]...)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(todos)
}
