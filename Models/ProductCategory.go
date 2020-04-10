package Models

import "github.com/jinzhu/gorm"

type ProductCategory struct {
	gorm.Model
	ProductCategoryName string `gorm:"not null"`
	Products []Product       
}