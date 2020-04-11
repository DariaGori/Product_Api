package Controllers

import (
	iv "github.com/DariaGori/Product_Api/Validation"
	s "github.com/DariaGori/Product_Api/Services"
	"github.com/DariaGori/Product_Api/Mappers"
	"github.com/DariaGori/Product_Api/DTOs"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductCategoryController struct {
	ProductCategoryService s.ProductCategoryService
	ProductService s.ProductService
}

func ProvideProductCategoryController(p s.ProductCategoryService, ps s.ProductService) ProductCategoryController {
	return ProductCategoryController{ProductCategoryService: p, ProductService: ps} 
}

// GET: ./
func (p *ProductCategoryController) FindAll(ctx *gin.Context) {
	productCategories := p.ProductCategoryService.FindAll()

	ctx.JSON(http.StatusOK, gin.H{"productCategories": Mappers.ToProductCategoryDTOs(productCategories)})
}

// GET: ./details/:id
func (p *ProductCategoryController) FindByID(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory, err := p.ProductCategoryService.FindByID(uint(id))

	// If no product category found, return bad request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": id})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"productCategory": Mappers.ToProductCategoryDTO(productCategory)})
}

// POST: ./create
func (p *ProductCategoryController) Create(ctx *gin.Context) {
	var productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO
	err := ctx.BindJSON(&productCategoryCreateEditDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	// Check if category with this name exists
	if p.ProductCategoryService.Exists(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category already exists": productCategoryCreateEditDTO.ProductCategoryName, "Created": false})
		return
	}

	// Do not allow empty or fully numeric category names
	if productCategoryCreateEditDTO.ProductCategoryName == "" || iv.IsNumeric(productCategoryCreateEditDTO.ProductCategoryName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category name isn't valid": productCategoryCreateEditDTO.ProductCategoryName, "Created": false})
		return
	}

	createdProductCategory := p.ProductCategoryService.Save(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO))

	ctx.JSON(http.StatusOK, gin.H{"Category created": createdProductCategory.ProductCategoryName})
}

// PUT: ./edit/:id
func (p *ProductCategoryController) Update(ctx *gin.Context) {
	var productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO
	err := ctx.BindJSON(&productCategoryCreateEditDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory, err := p.ProductCategoryService.FindByID(uint(id))

	// If no product category found, return bad request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": id})
		return
	}

	// Check if category with this name exists
	if p.ProductCategoryService.Exists(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category already exists": productCategoryCreateEditDTO.ProductCategoryName, "Updated": false})
		return
	}

	// Do not allow empty or fully numeric category names
	if productCategoryCreateEditDTO.ProductCategoryName == "" || iv.IsNumeric(productCategoryCreateEditDTO.ProductCategoryName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Category name is invalid": productCategoryCreateEditDTO.ProductCategoryName, "Created": false})
		return
	}

	productCategory.ProductCategoryName = productCategoryCreateEditDTO.ProductCategoryName
	p.ProductCategoryService.Save(productCategory)

	ctx.Status(http.StatusOK)
}

// DELETE: ./delete/:id
func (p *ProductCategoryController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory, err := p.ProductCategoryService.FindByID(uint(id))

	// If no product category found, return bad request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": id, "Deleted": false})
		return
	}

	// Delete products that belong to the category
	products := p.ProductService.FindByCategoryId(uint(id))
	for i := 0; i < len(products); i++ {
		p.ProductService.Delete(products[i])
	}

	p.ProductCategoryService.Delete(productCategory)

	ctx.Status(http.StatusOK)
}