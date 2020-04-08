package Services

import (
	"github.com/DariaGori/Product_Api/Repositories"
	"github.com/DariaGori/Product_Api/Models"
)

type ProductService struct {
	ProductRepository Repositories.ProductRepository
}

func ProvideProductService(p Repositories.ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() []Models.Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindByCategoryId(categoryId uint) []Models.Product {
	return p.ProductRepository.FindByCategoryId(categoryId)
}

func (p *ProductService) FindByID(id uint) Models.Product {
	return p.ProductRepository.FindByID(id)
}

func (p *ProductService) Save(product Models.Product) Models.Product {
	p.ProductRepository.Save(product)

	return product
}

func (p *ProductService) Delete(product Models.Product) {
	p.ProductRepository.Delete(product)
}