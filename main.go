package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var Users = []User{
	{ID: 1, UserName: "Beka", Password: "1234"},
	{ID: 2, UserName: "Nurs", Password: "1234"},
	{ID: 3, UserName: "Uluk", Password: "1234"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8000")
}
