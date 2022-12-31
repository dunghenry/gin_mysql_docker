package controllers

import (
	"gin/mysql/configs"
	"gin/mysql/models"
	"log"
	"net/http"
	"strconv"

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

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	row, err := configs.GetDataBase().Query("SELECT  id, name, email, mobile, password FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email, &user.Mobile, &user.Password)
	}
	defer row.Close()
	if user.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"msg": "User not found",
	})

}

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	stmt, err := configs.GetDataBase().Prepare("INSERT INTO users(name, email, mobile, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err,
		})
	}
	rs, err := stmt.Exec(user.Name, user.Email, user.Mobile, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	user.Id = int(id)
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	stmt, err := configs.GetDataBase().Prepare("UPDATE users SET name=?, email=?, mobile=?, password=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	rs, err := stmt.Exec(user.Name, user.Email, user.Mobile, user.Password, id)
	if err != nil {
		log.Fatal(err)
	}

	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	i, _ := strconv.Atoi(id)
	user.Id = i
	defer stmt.Close()
	if row == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": true,
			"message": "User not  found or update failed",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"user":    user,
		})
	}

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	stmt, err := configs.GetDataBase().Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	rs, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows := row
	if row == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "User not  found or deleted failed",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"count":   rows,
		})
	}

}
