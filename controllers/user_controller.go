package controllers

import (
	"net/http"

	"github.com/YahyaCengiz/go_deneme/CRUD/services"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    role := c.MustGet("role").(string)

    if role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"message": "You don't have permission to access this resource"})
        return
    }

    users, err := services.LoadUsers("./users.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}
