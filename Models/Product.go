package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName string `gorm:"not null"`
	ProductCategoryID uint
}