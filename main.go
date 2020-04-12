package main

import (
	"github.com/DariaGori/Product_Api/Models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := Models.SetupModels()

	r := SetupRouter(db)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

