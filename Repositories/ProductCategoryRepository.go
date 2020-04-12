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

// Return all product categories
func (p *ProductCategoryRepository) FindAll() []Models.ProductCategory {
	var productCategories []Models.ProductCategory
	var products []Models.Product
	p.DB.Find(&productCategories)
	// Get products of product category
	for i := 0; i < len(productCategories); i++ {
		p.DB.Where("products.product_category_id = ?", productCategories[i].ID).Find(&products)
		productCategories[i].Products = products
	}

	return productCategories
}

// Find product category by id
func (p *ProductCategoryRepository) FindByID(id uint) (Models.ProductCategory, error) {
	var productCategory Models.ProductCategory
	var products []Models.Product
	// If none found, return error
	if err := p.DB.First(&productCategory, id).Error; err != nil {
		return productCategory, err
	}
	// If found, get array of products
	p.DB.Where("products.product_category_id = ?", productCategory.ID).Find(&products)
	productCategory.Products = products
	return productCategory, nil
}

// Save product category to DB
func (p *ProductCategoryRepository) Save(productCategory Models.ProductCategory) Models.ProductCategory {
	p.DB.Save(&productCategory)
	return productCategory
}

// Delete product category
func (p *ProductCategoryRepository) Delete(productCategory Models.ProductCategory) {
	p.DB.Delete(&productCategory)
}

// Check if category with the name provided exists and isn't deleted by current time
func (p *ProductCategoryRepository) Exists(newCategory Models.ProductCategory) bool {
	var productCategory Models.ProductCategory
	if err := p.DB.Where("product_categories.product_category_name = ?", newCategory.ProductCategoryName).First(&productCategory).Error; err != nil {
		return false
	}

	return true
}

