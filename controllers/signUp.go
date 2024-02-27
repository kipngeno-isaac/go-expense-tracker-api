package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-expense-tracker-api/config"
	"github.com/kipngeno-isaac/go-expense-tracker-api/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignUp(c *gin.Context) {
	// get raw data
	var signupData struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}

	c.Bind(&signupData)
	// validate data

	// encrypt password
	hashedPassword, _ := HashPassword(signupData.Password)

	// create user
	user := models.User{
		FirstName: signupData.FirstName,
		LastName:  signupData.LastName,
		Email:     signupData.Email,
		Password:  hashedPassword,
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return token

	c.JSON(200, gin.H{"user": &user})
}
