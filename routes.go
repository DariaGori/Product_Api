package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	productController := InitProductController(db)
	productCategoryController := InitProductCategoryController(db)

	products := r.Group("/products")
	{
		products.GET("index", productController.FindAll)
		products.POST("create", productController.Create)
		products.GET("details/:productId", productController.FindByID)
		products.PUT("edit/:productId", productController.Update)
		products.DELETE("delete/:productId", productController.Delete)
	}
	productCategories := r.Group("/product-categories")
	{
		productCategories.GET("index", productCategoryController.FindAll)
		productCategories.POST("create", productCategoryController.Create)
		productCategories.GET("details/:productCategoryId", productCategoryController.FindByID)
		productCategories.PUT("edit/:productCategoryId", productCategoryController.Update)
		productCategories.DELETE("delete/:productCategoryId", productCategoryController.Delete)
	}
	return r
}