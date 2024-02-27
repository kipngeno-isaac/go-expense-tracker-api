package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-expense-tracker-api/config"
	"github.com/kipngeno-isaac/go-expense-tracker-api/controllers"
)

func init() {
	config.ConnectToDB()
	config.Migrate()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("sign-up", controllers.SignUp)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
