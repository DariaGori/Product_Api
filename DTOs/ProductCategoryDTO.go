package DTOs

type ProductCategoryDTO struct {
	Id uint 	               `json:"id" gorm:"primary_key";"AUTO_INCREMENT"`
	ProductCategoryName string `json:"productCategoryName"`
}