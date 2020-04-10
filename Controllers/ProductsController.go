package Controllers

import (
	"github.com/DariaGori/Product_Api/Services"
	"github.com/DariaGori/Product_Api/Mappers"
	"github.com/DariaGori/Product_Api/DTOs"
	"strconv"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService Services.ProductService
}

func ProvideProductController(p Services.ProductService) ProductController {
	return ProductController{ProductService: p}
}

func (p *ProductController) FindAll(ctx *gin.Context) {
	products := p.ProductService.FindAll()

	ctx.JSON(http.StatusOK, gin.H{"products": Mappers.ToProductDTOs(products)})
}

func (p *ProductController) FindByCategoryId(ctx *gin.Context) {
	categoryId, _ :=  strconv.Atoi(ctx.Param("productCategoryId"))
	products := p.ProductService.FindByCategoryId(uint(categoryId))
	
	ctx.JSON(http.StatusOK, gin.H{"products": Mappers.ToProductDTOs(products)})
}

func (p *ProductController) FindByID(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	
	ctx.JSON(http.StatusOK, gin.H{"product": Mappers.ToProductDTO(product)})
}

func (p *ProductController) Create(ctx *gin.Context) {
	var productDTO DTOs.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	if (productDTO.ProductName == "") {
		ctx.Status(http.StatusBadRequest)
		return
	}

	createdProduct := p.ProductService.Save(Mappers.ToProduct(productDTO))

	ctx.JSON(http.StatusOK, gin.H{"product": Mappers.ToProductDTO(createdProduct)})
}

func (p *ProductController) Update(ctx *gin.Context) {
	var productDTO DTOs.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product.ProductName == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	product.ProductName = productDTO.ProductName
	product.ProductCategoryID = productDTO.ProductCategoryID
	p.ProductService.Save(product)

	ctx.Status(http.StatusOK)
}

func (p *ProductController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product := p.ProductService.FindByID(uint(id))
	if product.ProductName == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	ctx.Status(http.StatusOK)
}