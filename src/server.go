package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Id   int
}

var users map[string]User

func useRESTServer() {
	r := gin.Default()
	users = make(map[string]User)
	count := 0

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/store", func(c *gin.Context) {
		name := c.Query("name")
		if _, ok := users[name]; !ok {
			user := User{Name: name, Id: count}
			users[name] = user
			c.JSON(http.StatusOK, gin.H{"message": "user created", "id": count})
			count++
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "user exists"})
		}
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})
	r.Run()
}

func main() {
	useRESTServer()
}
