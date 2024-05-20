package main

import (
	"github.com/YahyaCengiz/go_deneme/CRUD/controllers"
	"github.com/YahyaCengiz/go_deneme/CRUD/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")

    r.POST("/login", controllers.Login)

    authorized := r.Group("/")
    authorized.Use(middlewares.AuthMiddleware())
    {
        authorized.GET("/users", controllers.GetUsers)

        authorized.GET("/todos", controllers.GetTodos)
        authorized.GET("/todos/:todoId", controllers.GetTodoByID)
        authorized.GET("/todos/:todoId/messages/:messageId", controllers.GetTodoMessageByID)
        authorized.POST("/todos", controllers.CreateTodoList)
        authorized.POST("/todos/:todoId/messages", controllers.CreateTodoMessage)
        authorized.PUT("/todos/:todoId", controllers.UpdateTodo)
        authorized.PUT("/todos/:todoId/messages/:messageId", controllers.UpdateTodoMessage)
        authorized.DELETE("/todos/:todoId", controllers.DeleteTodo)
        authorized.DELETE("/todos/:todoId/messages/:messageId", controllers.DeleteTodoMessage)
    }

    r.Run(":3000")
}
