package Services

import (
	"github.com/DariaGori/Product_Api/Repositories"
	"github.com/DariaGori/Product_Api/Models"
)

type ProductCategoryService struct {
	ProductCategoryRepository Repositories.ProductCategoryRepository
}

func ProvideProductCategoryService(p Repositories.ProductCategoryRepository) ProductCategoryService {
	return ProductCategoryService{Repositories.ProductCategoryRepository: p}
}

func (p *ProductCategoryService) FindAll() []Models.ProductCategory {
	return p.ProductCategoryRepository.FindAll()
}

func (p *ProductCategoryService) FindByID(id uint) Models.ProductCategory {
	return p.ProductCategoryRepository.FindByID(id)
}

func (p *ProductCategoryService) Save(productCategory Models.ProductCategory) Models.ProductCategory {
	p.ProductRepository.Save(productCategory)

	return productCategory
}

func (p *ProductCategoryService) Delete(productCategory Models.ProductCategory) {
	p.ProductCategoryRepository.Delete(productCategory)
}