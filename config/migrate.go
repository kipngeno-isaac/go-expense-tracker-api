package config

import "github.com/kipngeno-isaac/go-expense-tracker-api/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Category{})
}
