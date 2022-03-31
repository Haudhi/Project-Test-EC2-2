package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name; type:text;" json:"name" form:"name"`
	Email    string `gorm:"column:email; type:text;" json:"email" form:"email"`
	Password string `gorm:"column:password; type:text;" json:"password" form:"password"`
	Rent     []Rent   `gorm:"foreignKey:UserID;reference:ID"`
	Book     []Book   `gorm:"foreignKey:UserID;reference:ID"`
}
