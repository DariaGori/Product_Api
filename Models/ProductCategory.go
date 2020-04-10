package Models

import "github.com/jinzhu/gorm"

type ProductCategory struct {
	gorm.Model
	ProductCategoryName string `gorm:"unique; not null"`
	Products []Product `gorm:"ForeignKey:ProductCategoryID" json:"products"`         
}