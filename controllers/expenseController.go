package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kipngeno-isaac/go-expense-tracker-api/config"
	"github.com/kipngeno-isaac/go-expense-tracker-api/models"
)

func SaveExpense(c *gin.Context) {
	// get expanes raw data
	var expenseData struct {
		UserId      int
		Name        string
		Amount      float64
		CategoryId  int
		Description string
	}

	c.Bind(&expenseData)

	// save expense
	expense := models.Expense{
		Name:        expenseData.Name,
		Amount:      expenseData.Amount,
		Description: expenseData.Description,
		CategoryId:  expenseData.CategoryId,
		UserId:      expenseData.UserId,
	}

	result := config.DB.Create(&expense)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"expense": &expense})
}
