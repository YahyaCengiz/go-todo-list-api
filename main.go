package main

import (
	"github.com/YahyaCengiz/go_deneme/CRUD/controllers"
	"github.com/YahyaCengiz/go_deneme/CRUD/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")

    // Public routes
    r.POST("/login", controllers.Login)

    // Protected routes
    authorized := r.Group("/")
    authorized.Use(middlewares.AuthMiddleware())
    {
        authorized.GET("/users", controllers.GetUsers)

        // Todos routes
        authorized.GET("/todos", controllers.GetTodos)
        authorized.GET("/todos/:id", controllers.GetTodoByID)
        authorized.POST("/todos", controllers.CreateTodoList)
        authorized.POST("/todos/:id/messages", controllers.CreateTodoMessage)
        authorized.PUT("/todos/:id", controllers.UpdateTodo)
        authorized.DELETE("/todos/:id", controllers.DeleteTodo)
    }

    r.Run(":3000")
}
