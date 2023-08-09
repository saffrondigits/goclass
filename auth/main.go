package main

import (
	"auth_service/db"
	"auth_service/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", loginHandler)
	router.POST("/register", registerHandler)

	fmt.Println("The auth server is running")
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

func loginHandler(c *gin.Context) {

}

func registerHandler(c *gin.Context) {
	var authReq models.AuthReq

	err := c.BindJSON(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request")
		return
	}

	client, err := db.DBConn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "db connection failed")
		return
	}

	var count int
	err = client.QueryRow("SELECT COUNT(*) FROM users where USERNAME = $1", authReq.Username).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database-error")
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, "Username already exists")
		return
	}

	_, err = client.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", authReq.Username, authReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "database-error")
		return
	}

	c.JSON(http.StatusCreated, "Registered successfully")
}
