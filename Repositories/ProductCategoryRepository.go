package Repositories

import (
	"github.com/DariaGori/Product_Api/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductCategoryRepository struct {
	DB *gorm.DB
}

func ProvideProductCategoryRepostiory(DB *gorm.DB) ProductCategoryRepository {
	return ProductCategoryRepository{DB: DB}
}

func (p *ProductCategoryRepository) FindAll() []Models.ProductCategory {
	var productCategories []Models.ProductCategory
	p.DB.Find(&productCategories)

	return productCategories
}

func (p *ProductCategoryRepository) FindByID(id uint) Models.ProductCategory {
	var productCategory Models.ProductCategory
	p.DB.First(&productCategory, id)

	return productCategory
}

func (p *ProductCategoryRepository) Save(productCategory Models.ProductCategory) Models.ProductCategory {
	p.DB.Save(&productCategory)

	return productCategory
}

func (p *ProductCategoryRepository) Delete(productCategory Models.ProductCategory) {
	p.DB.Delete(&productCategory)
}