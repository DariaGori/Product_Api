package Models

import "github.com/jinzhu/gorm"

type ProductCategory struct {
	gorm.Model
	ProductCategoryName string
	Products []Product `gorm:"ForeignKey:ProductCategoryID" json:"products"`         
}