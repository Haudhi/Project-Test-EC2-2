package entities

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	UserID uint `json:"UserID" form:"UserID"`
	BookID uint `json:"BookID" form:"BookID"`
	ReturnDate time.Time `json:"return" form:"return"`
	Status string `gorm:"default:rented" json:"status" form:"status"`
	Address string `json:"address" form:"address"`
}