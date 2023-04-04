package routes

import (
	"dumbmerch/handlers"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group) {
	e.GET("/todos", handlers.FindTodos)
	e.GET("/todos/:id", handlers.GetTodo)
	e.POST("/todos", handlers.CreateTodo)
	e.PATCH("/todos/:id", handlers.UpdateTodo)
	e.DELETE("/todos/:id", handlers.DeleteTodo)
}
