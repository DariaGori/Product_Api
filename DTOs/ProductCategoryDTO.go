package DTOs

type ProductCategoryDTO struct {
	Id uint 	               `json:"id, omitempty" gorm:"primary_key";"AUTO_INCREMENT"`
	ProductCategoryName string `json:"productCategoryName" gorm:"unique; not null"`
}