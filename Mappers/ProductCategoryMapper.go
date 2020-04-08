package Mappers

import (
	"github.com/DariaGori/Product_Api/DTOs"
	"github.com/DariaGori/Product_Api/Models"
)

func ToProductCategory(productCategoryDTO DTOs.ProductCategoryDTO) Models.ProductCategory {
	return Models.ProductCategory{ProductCategoryName: productCategoryDTO.ProductCategoryName}
}

func ToProductCategoryDTO(productCategory Models.ProductCategory) DTOs.ProductCategoryDTO {
	return ProductCategoryDTO{Id: productCategory.Id, ProductCategoryName: productCategory.ProductCategoryName}
}

func ToProductCategoryDTOs(productCategories []Models.ProductCategory) []DTOs.ProductCategoryDTO {
	productCategoryDtos := make([]DTOs.ProductCategoryDTO, len(productCategories))

	for i, itm := range productCategories {
		productCategoryDtos[i] = ToProductCategoryDTO(itm)
	}

	return productCategoryDtos
}