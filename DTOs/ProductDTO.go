package DTOs

type ProductDTO struct {
	Id		uint 	`json:"id, string, omitempty"`
	ProductName string     `json:"product_name"`
    ProductCategoryID uint `json:"product_category_id, string"`
}
