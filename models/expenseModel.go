package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Name        string
	Amount      float64
	Description string
	CategoryId  int
	UserId      int
}
