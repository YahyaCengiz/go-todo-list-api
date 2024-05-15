package controllers

import (
	"net/http"

	"github.com/YahyaCengiz/go_deneme/CRUD/models"
	"github.com/YahyaCengiz/go_deneme/CRUD/services"
	"github.com/gin-gonic/gin"
)

type PageData struct {
	Users []models.User
}

func GetUsers(c *gin.Context) {
    users, err := services.LoadUsers("./users.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load users data"})
        return
    }
    data := PageData{Users: users}

    c.HTML(http.StatusOK, "user_list.html", data)
}
