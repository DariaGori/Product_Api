package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	productController := InitProductController(db)
	productCategoryController := InitProductCategoryController(db)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "HEAD", "DELETE"}

	r.Use(cors.New(config))

	products := r.Group("/products")
	{
		products.GET("index", productController.FindAll)
		products.POST("create", productController.Create)
		products.GET("details/:id", productController.FindByID)
		products.PUT("edit/:id", productController.Update)
		products.DELETE("delete/:id", productController.Delete)
	}
	productCategories := r.Group("/product-categories")
	{
		productCategories.GET("index", productCategoryController.FindAll)
		productCategories.POST("create", productCategoryController.Create)
		productCategories.GET("details/:id", productCategoryController.FindByID)
		productCategories.PUT("edit/:id", productCategoryController.Update)
		productCategories.DELETE("delete/:id", productCategoryController.Delete)
	}
	return r
}