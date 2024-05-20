package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/YahyaCengiz/go_deneme/CRUD/models"
	"github.com/gin-gonic/gin"
)

var todos []models.Todo

func init() {
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

    var visibleTodos []models.Todo
    for _, todo := range todos {
        if todo.DeletedAt == nil && (role == "admin" || todo.UserID == username) {
            var visibleMessages []models.Message
            for _, message := range todo.Messages {
                if message.DeletedAt == nil {
                    visibleMessages = append(visibleMessages, message)
                }
            }
            todo.Messages = visibleMessages

            var completedMessages int
            for _, message := range todo.Messages {
                if message.IsCompleted {
                    completedMessages++
                }
            }
            if len(todo.Messages) > 0 {
                completionPercentage := (completedMessages * 100) / len(todo.Messages)
                todo.CompletionPercentage = completionPercentage
            }

            visibleTodos = append(visibleTodos, todo)
        }
    }

    c.JSON(http.StatusOK, visibleTodos)
}


func GetTodoByID(c *gin.Context) {
    id := c.Param("todoId")
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for _, todo := range todos {
        if todo.ID == id && todo.DeletedAt == nil {
            if role == "admin" || todo.UserID == username {
                var visibleMessages []models.Message
                for _, message := range todo.Messages {
                    if message.DeletedAt == nil {
                        visibleMessages = append(visibleMessages, message)
                    }
                }
                todo.Messages = visibleMessages
                c.JSON(http.StatusOK, todo)
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to view this todo"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}


func GetTodoMessageByID(c *gin.Context) {
    todoID := c.Param("todoId")
    messageID := c.Param("messageId")
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for _, todo := range todos {
        if todo.ID == todoID && todo.DeletedAt == nil {
            if role == "admin" || todo.UserID == username {
                for _, message := range todo.Messages {
                    if message.ID == messageID && message.DeletedAt == nil {
                        c.JSON(http.StatusOK, message)
                        return
                    }
                }
                c.JSON(http.StatusNotFound, gin.H{"message": "Message not found"})
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to view this message"})
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
    
    username, exists := c.Get("username")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing username in JWT token"})
        return
    }

    newTodo.UserID = username.(string)
    newTodo.ID = generateTodoID()
    newTodo.CreatedAt = time.Now()
    newTodo.UpdatedAt = time.Now()
    newTodo.CompletionPercentage = 0
    newTodo.Messages = []models.Message{}
    
    todos = append(todos, newTodo)
    saveTodos()
    c.JSON(http.StatusCreated, newTodo)
}


func CreateTodoMessage(c *gin.Context) {
    todoID := c.Param("todoId")
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
                newMessage.ID = generateMessageID(todo.Messages)
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
                updatedTodo.UserID = todo.UserID 
                updatedTodo.CreatedAt = todo.CreatedAt 
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

func UpdateTodoMessage(c *gin.Context) {
    todoID := c.Param("todoId")
    messageID := c.Param("messageId")
    var updatedMessage models.Message
    if err := c.ShouldBindJSON(&updatedMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for i, todo := range todos {
        if todo.ID == todoID {
            if role == "admin" || todo.UserID == username {
                for j, message := range todo.Messages {
                    if message.ID == messageID {
                        updatedMessage.ID = message.ID 
                        updatedMessage.CreatedAt = message.CreatedAt 
                        updatedMessage.UpdatedAt = time.Now()
                        todo.Messages[j] = updatedMessage
                        todo.UpdatedAt = time.Now()
                        todos[i] = todo
                        saveTodos()
                        c.JSON(http.StatusOK, updatedMessage)
                        return
                    }
                }
                c.JSON(http.StatusNotFound, gin.H{"message": "Message not found"})
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to update this message"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func DeleteTodoMessage(c *gin.Context) {
    todoID := c.Param("todoId")
    messageID := c.Param("messageId")
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for i, todo := range todos {
        if todo.ID == todoID {
            if role == "admin" || todo.UserID == username {
                for j, message := range todo.Messages {
                    if message.ID == messageID {
                        now := time.Now()
                        todo.Messages[j].DeletedAt = &now
                        todo.UpdatedAt = now
                        todos[i] = todo
                        saveTodos()
                        c.JSON(http.StatusOK, gin.H{"message": "Message marked as deleted"})
                        return
                    }
                }
                c.JSON(http.StatusNotFound, gin.H{"message": "Message not found"})
                return
            }
            c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to delete this message"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("todoId")
    role := c.MustGet("role").(string)
    username := c.MustGet("username").(string)

    for i, todo := range todos {
        if todo.ID == id {
            if role == "admin" || todo.UserID == username {
                now := time.Now()
                todos[i].DeletedAt = &now
                saveTodos()
                c.JSON(http.StatusOK, gin.H{"message": "Todo marked as deleted"})
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

func generateTodoID() string {
    maxID := 0
    for _, todo := range todos {
        id, err := strconv.Atoi(todo.ID[4:])
        if err == nil && id > maxID {
            maxID = id
        }
    }
    return "todo" + strconv.Itoa(maxID+1)
}

func generateMessageID(messages []models.Message) string {
    maxID := 0
    for _, message := range messages {
        id, err := strconv.Atoi(message.ID[7:])
        if err == nil && id > maxID {
            maxID = id
        }
    }
    return "message" + strconv.Itoa(maxID+1)
}
