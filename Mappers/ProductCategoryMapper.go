package Mappers

import (
	"github.com/DariaGori/Product_Api/DTOs"
	"github.com/DariaGori/Product_Api/Models"
)

func ToProductCategory(productCategoryDTO DTOs.ProductCategoryDTO) Models.ProductCategory {
	return Models.ProductCategory{ProductCategoryName: productCategoryDTO.ProductCategoryName, Products: productCategoryDTO.Products}
}

func ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO) Models.ProductCategory {
	return Models.ProductCategory{ProductCategoryName: productCategoryCreateEditDTO.ProductCategoryName}
}

func ToProductCategoryDTO(productCategory Models.ProductCategory) DTOs.ProductCategoryDTO {
	return DTOs.ProductCategoryDTO{Id: productCategory.ID, ProductCategoryName: productCategory.ProductCategoryName, Products: productCategory.Products}
}

func ToProductCategoryCreateEditDTO(productCategory Models.ProductCategory) DTOs.ProductCategoryCreateEditDTO {
	return DTOs.ProductCategoryCreateEditDTO{Id: productCategory.ID, ProductCategoryName: productCategory.ProductCategoryName}
}

func ToProductCategoryDTOs(productCategories []Models.ProductCategory) []DTOs.ProductCategoryDTO {
	productCategoryDtos := make([]DTOs.ProductCategoryDTO, len(productCategories))

	for i, itm := range productCategories {
		productCategoryDtos[i] = ToProductCategoryDTO(itm)
	}

	return productCategoryDtos
}

func ToProductCategoryCreateEditDTOs(productCategories []Models.ProductCategory) []DTOs.ProductCategoryCreateEditDTO {
	productCategoryCreateEditDtos := make([]DTOs.ProductCategoryCreateEditDTO, len(productCategories))

	for i, itm := range productCategories {
		productCategoryCreateEditDtos[i] = ToProductCategoryCreateEditDTO(itm)
	}

	return productCategoryCreateEditDtos
}