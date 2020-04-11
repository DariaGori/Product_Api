package Repositories

import (
	"time"
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

// Return all products
func (p *ProductRepository) FindAll() []Models.Product {
	var products []Models.Product
	p.DB.Find(&products)

	return products
}

// Return all products with matching product category id
func (p *ProductRepository) FindByCategoryId(productCategoryId uint) []Models.Product {
	var products []Models.Product
	p.DB.Where("product_category_id = ?", productCategoryId).Find(&products)

	return products
}

// Find product by id
func (p *ProductRepository) FindByID(id uint) (Models.Product, error) {
	var product Models.Product
	// If none found, return error
	if err := p.DB.First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

// Save product to DB
func (p *ProductRepository) Save(product Models.Product) Models.Product {
	p.DB.Save(&product)

	return product
}

// Delete product
func (p *ProductRepository) Delete(product Models.Product) {
	p.DB.Delete(&product)
}

// Check if product with the name provided exists and isn't deleted by current time
func (p *ProductRepository) Exists(newProduct Models.Product) bool {
	var product Models.Product
	query := p.DB.Where("products.product_name = ?", newProduct.ProductName)
	if err := query.Where("products.deleted_at IS NULL OR products.deleted_at > ?", time.Now()).First(&product).Error; err != nil {
		return false
	}
	return true
}