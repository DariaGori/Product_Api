package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	productController := InitProductController(db)
	productCategoryController := InitProductCategoryController(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
        AllowHeaders:     []string{"Content-Type, Content-Length, Origin"},
        ExposeHeaders:    []string{"Content-Type, Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:8080"
        },
        MaxAge: 12 * time.Hour,
    }))

	productApi := r.Group("/product-api")
	{
		productApi.GET("products", productController.FindAll)
		productApi.GET("products/by-category/:category-id", productController.FindByCategoryId)
		productApi.POST("products/create", productController.Create)
		productApi.GET("products/details/:id", productController.FindByID)
		productApi.PUT("products/edit/:id", productController.Update)
		productApi.DELETE("products/delete/:id", productController.Delete)
		productApi.GET("product-categories", productCategoryController.FindAll)
		productApi.POST("product-categories/create", productCategoryController.Create)
		productApi.GET("product-categories/details/:id", productCategoryController.FindByID)
		productApi.PUT("product-categories/edit/:id", productCategoryController.Update)
		productApi.DELETE("product-categories/delete/:id", productCategoryController.Delete)
	}
	return r
}