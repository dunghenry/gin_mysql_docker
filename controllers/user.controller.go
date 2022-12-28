package controllers

import (
	"gin/mysql/configs"
	"gin/mysql/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := configs.GetDataBase().Query("SELECT id, name, email, mobile, password FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.Mobile, &user.Password)
		users = append(users, user)
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": users,
		"count":  len(users),
	})
}
