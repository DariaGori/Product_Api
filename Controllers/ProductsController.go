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

type ProductController struct {
	ProductService s.ProductService
	ProductCategoryService s.ProductCategoryService
}

func ProvideProductController(p s.ProductService, pc s.ProductCategoryService) ProductController {
	return ProductController{ProductService: p, ProductCategoryService: pc}
}

// GET: ./
func (p *ProductController) FindAll(ctx *gin.Context) {
	products := p.ProductService.FindAll()

	ctx.JSON(http.StatusOK, gin.H{"products": Mappers.ToProductDTOs(products)})
}

// GET: ./by-category/:category-id
func (p *ProductController) FindByCategoryId(ctx *gin.Context) {
	categoryId, _ :=  strconv.Atoi(ctx.Param("category-id"))
	products := p.ProductService.FindByCategoryId(uint(categoryId))
	
	ctx.JSON(http.StatusOK, gin.H{"products": Mappers.ToProductDTOs(products)})
}

// GET: ./details/:id
func (p *ProductController) FindByID(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))

	// If no product found, return bad request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No product found by this id": id})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"product": Mappers.ToProductDTO(product)})
}

// POST: ./create
func (p *ProductController) Create(ctx *gin.Context) {
	var productDTO DTOs.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	// Check if product with this name exists
	if p.ProductService.Exists(Mappers.ToProduct(productDTO)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Product already exists": productDTO.ProductName, "Created": false})
		return
	}

	// Do not allow empty or fully numeric product names
	if productDTO.ProductName == "" || iv.IsNumeric(productDTO.ProductName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Product name isn't valid": productDTO.ProductName, "Created": false})
		return
	}

	// If no category is found by category id
	productCategory, err := p.ProductCategoryService.FindByID(productDTO.ProductCategoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": productDTO.ProductCategoryID, "Created": false})
		return
	}

	// If category found is deleted
	if !(p.ProductCategoryService.Exists(productCategory)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": productDTO.ProductCategoryID, "Created": false})
		return
	}

	createdProduct := p.ProductService.Save(Mappers.ToProduct(productDTO))

	ctx.JSON(http.StatusOK, gin.H{"Product created": createdProduct.ProductName})
}

// PUT: ./edit/:id
func (p *ProductController) Update(ctx *gin.Context) {
	var productDTO DTOs.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))

	// If no product found, return bad request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No product found by this id": id})
		return
	}

	// Check if product with this name exists
	if p.ProductService.Exists(product) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Product already exists": productDTO.ProductName, "Updated": false})
		return
	}

	// Do not allow empty or fully numeric product names
	if productDTO.ProductName == "" || iv.IsNumeric(productDTO.ProductName) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Product name is invalid": productDTO.ProductName, "Created": false})
		return
	}

	// If no category is found by category id
	productCategory, err := p.ProductCategoryService.FindByID(productDTO.ProductCategoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": productDTO.ProductCategoryID, "Created": false})
		return
	}

	// If category found is deleted
	if !(p.ProductCategoryService.Exists(productCategory)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"No category found by this id": productDTO.ProductCategoryID, "Created": false})
		return
	}

	product.ProductName = productDTO.ProductName
	product.ProductCategoryID = productDTO.ProductCategoryID
	p.ProductService.Save(product)

	ctx.Status(http.StatusOK)
}

// DELETE: ./delete/:id
func (p *ProductController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))

	// If no product found, return bad request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No product found by this id": id, "Deleted": false})
		return
	}

	p.ProductService.Delete(product)

	ctx.Status(http.StatusOK)
}