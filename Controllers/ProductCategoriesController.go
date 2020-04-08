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

type ProductCategoryController struct {
	ProductCategoryService Services.ProductCategoryService
}

func ProvideProductCategoryController(p Services.ProductCategoryService) ProductCategoryController {
	return ProductCategoryController{Services.ProductCategoryService: p}
}

func (p *ProductCategoryController) FindAll(ctx *gin.Context) {
	productCategories := p.ProductCategoryService.FindAll()

	ctx.JSON(http.StatusOK, gin.H{"productCategories": Mappers.ToProductCategoryDTOs(productCategories)})
}

func (p *ProductCategoryController) FindByID(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory := p.ProductCategoryService.FindByID(uint(id))
	
	ctx.JSON(http.StatusOK, gin.H{"productCategory": Mappers.ToProductCategoryDTO(productCategory)})
}

func (p *ProductCategoryController) Create(ctx *gin.Context) {
	var productCategoryDTO DTOs.ProductCategoryDTO
	err := ctx.BindJSON(&productCategoryDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	createdProductCategory := p.ProductCategoryService.Save(Mappers.ToProductCategory(productCategoryDTO))

	ctx.JSON(http.StatusOK, gin.H{"productCategory": Mappers.ToProductCategoryDTO(createdProductCategory)})
}

func (p *ProductCategoryController) Update(ctx *gin.Context) {
	var productCategoryDTO DTOs.ProductCategoryDTO
	err := ctx.BindJSON(&productCategoryDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory := p.ProductCategoryService.FindByID(uint(id))
	if productCategory == (ProductCategory{}) {
		ctx.Status(http.StatusBadRequest)
		return
	}

	productCategory.ProductCategoryName = productCategoryDTO.ProductCategoryName
	p.ProductCategoryService.Save(productCategory)

	ctx.Status(http.StatusOK)
}

func (p *ProductCategoryController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory := p.ProductCategoryService.FindByID(uint(id))
	if productCategory == (ProductCategory{}) {
		ctx.Status(http.StatusBadRequest)
		return
	}

	p.ProductCategoryService.Delete(productCategory)

	ctx.Status(http.StatusOK)
}