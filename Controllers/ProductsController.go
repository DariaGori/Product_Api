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

	// If no product found, return not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Feedback":"No product found by this id: " + strconv.Itoa(id), "Status":http.StatusNotFound})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{"product": Mappers.ToProductDTO(product)})
}

// POST: ./create
func (p *ProductController) Create(ctx *gin.Context) {
	var productDTO DTOs.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"JSON decoding: " + err.Error(), "Status": http.StatusBadRequest})
		return
	}

	// Check if product with this name exists
	if p.ProductService.Exists(Mappers.ToProduct(productDTO)) {
		ctx.JSON(http.StatusConflict, gin.H{"Feedback":"Product name already exists: " + productDTO.ProductName + ". Product isn't created", "Status": http.StatusConflict})
		return
	}

	// Do not allow fully numeric product names
	if iv.IsNumeric(productDTO.ProductName) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"Feedback":"Product name isn't valid: " + productDTO.ProductName + ". Product isn't created", "Status": http.StatusNotAcceptable})
		return
	}

	// If no product is found by category id
	productCategory, err := p.ProductCategoryService.FindByID(productDTO.ProductCategoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(int(productDTO.ProductCategoryID)) + ". Product isn't created", "Status": http.StatusBadRequest})
		return
	}

	// If category found is deleted
	if !(p.ProductCategoryService.Exists(productCategory)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(int(productDTO.ProductCategoryID)) + ". Product isn't created", "Status": http.StatusBadRequest})
		return
	}

	createdProduct := p.ProductService.Save(Mappers.ToProduct(productDTO))

	ctx.JSON(http.StatusCreated, gin.H{"Feedback":"Product created: " + createdProduct.ProductName, "Status":http.StatusCreated})
}

// PUT: ./edit/:id
func (p *ProductController) Update(ctx *gin.Context) {
	var productDTO DTOs.ProductDTO
	err := ctx.BindJSON(&productDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"JSON decoding: " + err.Error(), "Status": http.StatusBadRequest})
		return
	}

	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))

	// If no product found, return not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Feedback":"No product found by this id: " + strconv.Itoa(id), "Status":http.StatusNotFound})
		return
	}

	// Check if product with this name exists
	if p.ProductService.Exists(product) {
		ctx.JSON(http.StatusConflict, gin.H{"Feedback":"Product name already exists: " + productDTO.ProductName + ". Product isn't updated", "Status": http.StatusConflict})
		return
	}

	// Do not allow fully numeric product names
	if iv.IsNumeric(productDTO.ProductName) {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"Feedback":"Product name isn't valid: " + productDTO.ProductName + ". Product isn't updated", "Status": http.StatusNotAcceptable})
		return
	}

	// If no category is found by category id
	productCategory, err := p.ProductCategoryService.FindByID(productDTO.ProductCategoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(int(productDTO.ProductCategoryID)) + ". Product isn't updated", "Status": http.StatusBadRequest})
		return
	}

	// If category found is deleted
	if !(p.ProductCategoryService.Exists(productCategory)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Feedback":"No category found by this id: " + strconv.Itoa(int(productDTO.ProductCategoryID)) + ". Product isn't updated", "Status": http.StatusBadRequest})
		return
	}

	product.ProductName = productDTO.ProductName
	product.ProductCategoryID = productDTO.ProductCategoryID
	p.ProductService.Save(product)

	ctx.JSON(http.StatusAccepted, gin.H{"Feedback":"Product updated: " + product.ProductName, "Status":http.StatusAccepted})
}

// DELETE: ./delete/:id
func (p *ProductController) Delete(ctx *gin.Context) {
	id, _ :=  strconv.Atoi(ctx.Param("id"))
	product, err := p.ProductService.FindByID(uint(id))

	// If no product found, return not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Feedback":"No product found by this id: " + strconv.Itoa(id), "Status":http.StatusNotFound})
		return
	}

	p.ProductService.Delete(product)

	ctx.JSON(http.StatusOK, gin.H{"Feedback":"Product deleted: " + product.ProductName, "Status":http.StatusOK})
}