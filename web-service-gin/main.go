package main

import (
	"errors"
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

type Todo struct {
	ID        string   `json:"id"`
	Label     string   `json:"label" binding:"required"`
	Completed bool     `json:"completed"`
	Priority  priority `json:"priority" binding:"required"`
}

type TodoCompletedReq struct {
	Completed *bool `json:"completed" binding:"required"`
}

var todos = []Todo{
	{ID: "-testing-", Label: "clean room", Completed: false, Priority: LOW},
	{ID: shortid.MustGenerate(), Label: "exercise", Completed: false, Priority: MEDIUM},
	{ID: shortid.MustGenerate(), Label: "learn Go", Completed: true, Priority: HIGH},
}

func main() {
	router := gin.Default()

	router.GET("/ping", healthCheck)
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodoByID)
	router.PUT("/todos/:id", updateTodo)

	router.Run()
}

func findTodoByID(id string) (*Todo, error) {
	for index, t := range todos {
		if t.ID == id {
			return &todos[index], nil
		}
	}

	return nil, errors.New("todo not found")
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
	var newTodo Todo

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
	todo, err := findTodoByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := findTodoByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	var body TodoCompletedReq
	if err := c.BindJSON(&body); err != nil {
		return
	}

	todo.Completed = *body.Completed
	c.IndentedJSON(http.StatusOK, todo)
}
