package main

import (
	"github.com/DariaGori/Product_Api/Models"
	//"os"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	db := Models.SetupModels()

	r := gin.Default()

	SetupRouter(r, db)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}