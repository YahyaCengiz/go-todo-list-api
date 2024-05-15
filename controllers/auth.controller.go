package controllers

import (
	"net/http"

	"github.com/YahyaCengiz/go_deneme/CRUD/services"
	"github.com/YahyaCengiz/go_deneme/CRUD/utils"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := services.AuthenticateUser(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := utils.GenerateJWT(user.Username, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
