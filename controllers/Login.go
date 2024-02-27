package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-expense-tracker-api/config"
	"github.com/kipngeno-isaac/go-expense-tracker-api/models"
)

func Login(c *gin.Context) {
	// get raw data
	var loginData struct {
		Email    string
		Password string
	}

	c.Bind(&loginData)
	// validate data

	// Find User
	var user models.User
	config.DB.First(&user, "email=?", loginData.Email)

	// verify password
	CheckPasswordHash(loginData.Password, user.Password)
	// return token

	c.JSON(200, gin.H{"user": &user})

}
