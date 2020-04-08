package Mappers

import (
	"github.com/DariaGori/Product_Api/DTOs"
	"github.com/DariaGori/Product_Api/Models"
)

func ToProduct(productDTO DTOs.ProductDTO) Models.Product {
	return Models.Product{ProductName: productDTO.ProductName, ProductCategoryID: productDTO.ProductCategoryID}
}

func ToProductDTO(product Models.Product) DTOs.ProductDTO {
	return ProductDTO{Id: product.Id, ProductName: product.ProductName, ProductCategoryID: product.ProductCategoryID}
}

func ToProductDTOs(products []Models.Product) []DTOs.ProductDTO {
	productDtos := make([]DTOs.ProductDTO, len(products))

	for i, itm := range products {
		productDtos[i] = ToProductDTO(itm)
	}

	return productDtos
}