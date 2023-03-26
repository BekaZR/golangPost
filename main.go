package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id"`
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

func addUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	Users = append(Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println("id ")
	if err != nil {
		return
	}
	for _, user := range Users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"detail": "User not found"})
}

func deleteUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println("id ")
	if err != nil {
		return
	}
	for _, user := range Users {
		if user.ID == id {
			Users = append(Users[:id], Users[id+1:]...)
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"detail": "User not found"})
}

func UpdateUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println("id ")
	if err != nil {
		return
	}
	var UpdateUser User
	if err := c.BindJSON(&UpdateUser); err != nil {
		return
	}
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == id {
			Users[i] = UpdateUser
			c.IndentedJSON(http.StatusOK, Users[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"detail": "User not found"})
}

func main() {

	router := gin.Default()
	router.GET("/user", getUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/user", addUser)
	router.DELETE("/user/:id", deleteUserById)
	router.PUT("/user/:id", UpdateUserById)
	router.Run("0.0.0.0:8000")
}
