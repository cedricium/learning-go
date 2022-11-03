package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

type priority int

const (
	LOW priority = iota
	MEDIUM
	HIGH
)

/* the type definition and constant above is a workaround for declaring enums */
/* references: https://go.dev/ref/spec#Iota */

type todo struct {
	ID        string   `json:"id"`
	Label     string   `json:"label" binding:"required"`
	Completed bool     `json:"completed"`
	Priority  priority `json:"priority" binding:"required"`
}

var todos = []todo{
	{ID: shortid.MustGenerate(), Label: "clean room", Completed: false, Priority: LOW},
	{ID: shortid.MustGenerate(), Label: "exercise", Completed: false, Priority: MEDIUM},
	{ID: shortid.MustGenerate(), Label: "learn Go", Completed: true, Priority: HIGH},
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.ID = shortid.MustGenerate()
	newTodo.Completed = false

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func updateTodo(c *gin.Context) {
	/* TODO */
}

func main() {
	router := gin.Default()

	router.GET("/ping", healthCheck)
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodoByID)
	router.PUT("/todos/:id", updateTodo)

	router.Run("localhost:8080")
}
