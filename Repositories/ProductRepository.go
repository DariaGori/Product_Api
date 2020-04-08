package Repositories

import (
	"github.com/DariaGori/Product_Api/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []Models.Product {
	var products []Models.Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByCategoryId(productCategoryId uint) []Models.Product {
	var products []Models.Product
	p.DB.Find(&products).Where("productCategoryId = ?", productCategoryId)

	return products
}

func (p *ProductRepository) FindByID(id uint) Models.Product {
	var product Models.Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product Models.Product) Models.Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product Models.Product) {
	p.DB.Delete(&product)
}