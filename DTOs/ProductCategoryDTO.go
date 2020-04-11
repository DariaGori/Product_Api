package DTOs

import (
	"github.com/DariaGori/Product_Api/Models"
)

type ProductCategoryDTO struct {
	Id uint 	               `json:"id" gorm:"primary_key";"AUTO_INCREMENT"`
	ProductCategoryName string `json:"productCategoryName" binding:"required"`
	Products []Models.Product  `json:"products" binding:"required"`
}