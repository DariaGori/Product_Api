package DTOs

type ProductDTO struct {
	Id uint 	             `json:"id, omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	ProductName string       `json:"product_name" gorm:"unique; not null"`
    ProductCategoryID uint   `json:"product_category_id"`
}
