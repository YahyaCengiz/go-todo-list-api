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
    }

    r.Run(":3000")
}
