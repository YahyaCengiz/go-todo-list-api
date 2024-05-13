package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/YahyaCengiz/go_deneme/CRUD/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnv()
}


type User struct{
	Name string `json:"name"`
	Role string `json:"role"`
	Age int `json:"age"`
}

func main() {
	router := gin.Default()


	router.GET("/users", getUsersHandler)

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	router.GET("/goodbye", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "goodbye",
		})
	})

	router.Run() // listen and serve on

}

func getUsersHandler(c *gin.Context) {
	usersData, err := ioutil.ReadFile("users.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []User
    if err := json.Unmarshal(usersData, &users); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse users data"})
        return
    }

	c.JSON(http.StatusOK, users)
}