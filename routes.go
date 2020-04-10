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
		productCategories.OPTIONS("details/:id", PreFlight)
		productCategories.PUT("edit/:id", productCategoryController.Update)
		productCategories.DELETE("delete/:id", productCategoryController.Delete)
	}
	return r
}