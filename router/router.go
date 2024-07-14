package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string
	Name string
	Age  int
}

var users = []User{
	{Id: "1", Name: "perry", Age: 28},
	{Id: "1", Name: "simon", Age: 29},
}

func StartServer() {
	server := gin.Default()
	server.GET("/users", getUsers)
	server.GET("/users/:id", getUserById)
	server.POST("/user/create", createUser)
	server.Run("localhost:8080")

}

func createUser(context *gin.Context) {
	var newUser User
	err := context.BindJSON(&newUser)
	if err != nil || newUser.Id == "" {
		return
	}
	users = append(users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)

}

func getUserById(context *gin.Context) {
	var id = context.Param("id")
	for _, user := range users {
		if user.Id == id {
			context.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}