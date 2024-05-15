package middlewares

import (
	"net/http"

	"github.com/YahyaCengiz/go_deneme/CRUD/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("username", claims.Username)
        c.Set("role", claims.Role)

        c.Next()
    }
}
