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
	return ProductCategoryController{ProductCategoryService: p}
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
	var productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO
	err := ctx.BindJSON(&productCategoryCreateEditDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Check if category with this name exists
	if p.ProductCategoryService.Exists(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category already exists": productCategoryCreateEditDTO.ProductCategoryName, "Created": false})
		return
	}

	if (productCategoryCreateEditDTO.ProductCategoryName == "") {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category name is empty": productCategoryCreateEditDTO.ProductCategoryName, "Created": false})
		return
	}

	createdProductCategory := p.ProductCategoryService.Save(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO))

	ctx.JSON(http.StatusOK, gin.H{"Category created": createdProductCategory.ProductCategoryName})
}

func (p *ProductCategoryController) Update(ctx *gin.Context) {
	var productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO
	err := ctx.BindJSON(&productCategoryCreateEditDTO)
	if err != nil {
		log.Fatalln(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// Check if category with this name exists
	if p.ProductCategoryService.Exists(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category already exists": productCategoryCreateEditDTO.ProductCategoryName, "Updated": false})
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory := p.ProductCategoryService.FindByID(uint(id))
	if productCategory.ProductCategoryName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category name is empty": productCategoryCreateEditDTO.ProductCategoryName, "Created": false})
		return
	}

	productCategory.ProductCategoryName = productCategoryCreateEditDTO.ProductCategoryName
	p.ProductCategoryService.Save(productCategory)

	ctx.Status(http.StatusOK)
}

func (p *ProductCategoryController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory := p.ProductCategoryService.FindByID(uint(id))
	if productCategory.ProductCategoryName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category id doesn't exist": id, "Deleted": false})
		return
	}

	p.ProductCategoryService.Delete(productCategory)

	ctx.Status(http.StatusOK)
}