package DTOs

type ProductDTO struct {
	Id uint 	             `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ProductName string       `json:"productName" gorm:"unique; not null" binding:"required"`
    ProductCategoryID uint   `json:"productCategoryId" binding:"required"`
}
