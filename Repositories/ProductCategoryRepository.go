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
	var products []Models.Product
	p.DB.Find(&productCategories)
	for i := 0; i < len(productCategories); i++ {
		p.DB.Where("products.product_category_id = ?", productCategories[i].ID).Find(&products)
		productCategories[i].Products = products
	}

	return productCategories
}

func (p *ProductCategoryRepository) FindByID(id uint) Models.ProductCategory {
	var productCategory Models.ProductCategory
	var products []Models.Product
	p.DB.Find(&productCategory)
	p.DB.Where("products.product_category_id = ?", productCategory.ID).Find(&products)
	productCategory.Products = products

	return productCategory
}

func (p *ProductCategoryRepository) Save(productCategory Models.ProductCategory) Models.ProductCategory {
	p.DB.Save(&productCategory)

	return productCategory
}

func (p *ProductCategoryRepository) Delete(productCategory Models.ProductCategory) {
	p.DB.Delete(&productCategory)
}

// Add check for deletedAt == null or less than current time
func (p *ProductCategoryRepository) Exists(newCategory Models.ProductCategory) bool {
	var productCategory Models.ProductCategory
	if err := p.DB.Where("product_categories.product_category_name = ?", newCategory.ProductCategoryName).First(&productCategory).Error; err != nil {
		return false
	}
	return true
}