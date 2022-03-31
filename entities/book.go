package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID     int `json:"user" form:"user"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Status string `gorm:"default:available" json:"status" form:"status"`
	Rent     []Rent   `gorm:"foreignKey:BookID;reference:ID"`

}