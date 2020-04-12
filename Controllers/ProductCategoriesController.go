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

	// If no product category found, return not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(id), "Status":http.StatusNotFound})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"productCategory": Mappers.ToProductCategoryDTO(productCategory)})
}

// POST: ./create
func (p *ProductCategoryController) Create(ctx *gin.Context) {
	var productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO
	err := ctx.BindJSON(&productCategoryCreateEditDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"JSON decoding: " + err.Error(), "Status": http.StatusBadRequest})
		return
	}

	// Check if category with this name exists
	if p.ProductCategoryService.Exists(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO)) {
		ctx.JSON(http.StatusConflict, gin.H{"Feedback":"Category name already exists: " + productCategoryCreateEditDTO.ProductCategoryName + ". Category isn't created", "Status": http.StatusConflict})
		return
	}

	// Do not allow fully numeric category names
	if iv.IsNumeric(productCategoryCreateEditDTO.ProductCategoryName) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"Feedback":"Category name isn't valid: " + productCategoryCreateEditDTO.ProductCategoryName + ". Category isn't created", "Status": http.StatusNotAcceptable})
		return
	}

	createdProductCategory := p.ProductCategoryService.Save(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO))

	ctx.JSON(http.StatusCreated, gin.H{"Feedback":"Category created: " + createdProductCategory.ProductCategoryName, "Status":http.StatusCreated})
}

// PUT: ./edit/:id
func (p *ProductCategoryController) Update(ctx *gin.Context) {
	var productCategoryCreateEditDTO DTOs.ProductCategoryCreateEditDTO
	err := ctx.BindJSON(&productCategoryCreateEditDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"JSON decoding: " + err.Error(), "Status": http.StatusBadRequest})
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory, err := p.ProductCategoryService.FindByID(uint(id))

	// If no product category found, return not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(id), "Status":http.StatusNotFound})
		return
	}

	// Check if category with this name exists
	if p.ProductCategoryService.Exists(Mappers.ToProductCategoryFromCreateEditDTO(productCategoryCreateEditDTO)) {
		ctx.JSON(http.StatusConflict, gin.H{"Feedback":"Category name already exists: " + productCategoryCreateEditDTO.ProductCategoryName + ". Category isn't updated", "Status": http.StatusConflict})
		return
	}

	// Do not allow fully numeric category names
	if iv.IsNumeric(productCategoryCreateEditDTO.ProductCategoryName) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"Feedback":"Category name isn't valid: " + productCategoryCreateEditDTO.ProductCategoryName + ". Category isn't updated", "Status": http.StatusNotAcceptable})
		return
	}

	productCategory.ProductCategoryName = productCategoryCreateEditDTO.ProductCategoryName
	p.ProductCategoryService.Save(productCategory)

	ctx.JSON(http.StatusAccepted, gin.H{"Feedback":"Category updated: " + productCategory.ProductCategoryName, "Status":http.StatusAccepted})
}

// DELETE: ./delete/:id
func (p *ProductCategoryController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	productCategory, err := p.ProductCategoryService.FindByID(uint(id))

	// If no product category found, return not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(id), "Status":http.StatusNotFound})
		return
	}

	// Delete products that belong to the category
	products := p.ProductService.FindByCategoryId(uint(id))
	for i := 0; i < len(products); i++ {
		p.ProductService.Delete(products[i])
	}

	p.ProductCategoryService.Delete(productCategory)

	ctx.JSON(http.StatusOK, gin.H{"Feedback":"Category deleted: " + productCategory.ProductCategoryName, "Status":http.StatusOK})
}