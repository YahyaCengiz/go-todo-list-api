package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/YahyaCengiz/go_deneme/CRUD/models"
	"github.com/gin-gonic/gin"
)


var todos []models.Todo

func init() {
    // Read todos from JSON file
    file, err := os.Open("todos.JSON")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    bytes, _ := ioutil.ReadAll(file)
    json.Unmarshal(bytes, &todos)
}

func GetTodos(c *gin.Context) {
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    if role == "admin" {
        c.JSON(http.StatusOK, todos)
    } else {
        var userTodos []models.Todo
        for _, todo := range todos {
            if todo.UserID == username {
                userTodos = append(userTodos, todo)
            }
        }
        c.JSON(http.StatusOK, userTodos)
    }
}

func GetTodoByID(c *gin.Context) {
    id := c.Param("id")
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for _, todo := range todos {
        if todo.ID == id {
            if role == "admin" || todo.UserID == username {
                c.JSON(http.StatusOK, todo)
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to view this todo"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func CreateTodoList(c *gin.Context) {
    var newTodo models.Todo
    if err := c.ShouldBindJSON(&newTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newTodo.UserID = c.MustGet("username").(string)
    newTodo.CreatedAt = time.Now()
    newTodo.UpdatedAt = time.Now()
    todos = append(todos, newTodo)
    saveTodos()
    c.JSON(http.StatusCreated, newTodo)
}

func CreateTodoMessage(c *gin.Context) {
    todoID := c.Param("id")
    var newMessage models.Message
    if err := c.ShouldBindJSON(&newMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for i, todo := range todos {
        if todo.ID == todoID {
            if role == "admin" || todo.UserID == username {
                newMessage.ID = generateID() // Function to generate a unique ID
                newMessage.CreatedAt = time.Now()
                newMessage.UpdatedAt = time.Now()
                todo.Messages = append(todo.Messages, newMessage)
                todo.UpdatedAt = time.Now()
                todos[i] = todo
                saveTodos()
                c.JSON(http.StatusCreated, newMessage)
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to add a message to this todo"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var updatedTodo models.Todo
    if err := c.ShouldBindJSON(&updatedTodo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for i, todo := range todos {
        if todo.ID == id {
            if role == "admin" || todo.UserID == username {
                updatedTodo.UserID = todo.UserID // Preserve original user ID
                updatedTodo.CreatedAt = todo.CreatedAt // Preserve original creation date
                updatedTodo.UpdatedAt = time.Now()
                todos[i] = updatedTodo
                saveTodos()
                c.JSON(http.StatusOK, updatedTodo)
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to update this todo"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for i, todo := range todos {
        if todo.ID == id {
            if role == "admin" || todo.UserID == username {
                todos = append(todos[:i], todos[i+1:]...)
                saveTodos()
                c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to delete this todo"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func saveTodos() {
    file, err := os.Create("todos.JSON")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    bytes, _ := json.MarshalIndent(todos, "", "  ")
    file.Write(bytes)
}

func generateID() string {
    // Implement a function to generate a unique ID
    return "new_unique_id"
}
