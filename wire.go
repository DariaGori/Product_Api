//+build wireinject

package main

import (
	"github.com/DariaGori/Product_Api/Controllers"
	"github.com/DariaGori/Product_Api/Repositories"
	"github.com/DariaGori/Product_Api/Services"
	"github.com/jinzhu/gorm"
	"github.com/google/wire"
)

func InitProductController(db *gorm.DB) Controllers.ProductController {
	wire.Build(Repositories.ProvideProductRepostiory, Services.ProvideProductService, Controllers.ProvideProductController)

	return Controllers.ProductController{}
}

func InitProductCategoryController(db *gorm.DB) Controllers.ProductCategoryController {
	wire.Build(Repositories.ProvideProductCategoryRepostiory, Services.ProvideProductCategoryService, Controllers.ProvideProductCategoryController)

	return Controllers.ProductCategoryController{}
}