package DTOs

type ProductCategoryDTO struct {
	ID uint 	               `json:"id, string, omitempty"`
	ProductCategoryName string `json:"product_category_name"`
}