package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductCategoryName string
	Products []Product         
}